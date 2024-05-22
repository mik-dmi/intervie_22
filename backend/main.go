package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type ContainerProperties struct {
	Id           string
	Names        []string
	Image        string
	Status       string
	CreationDate string
}

const listenAddr = "0.0.0.0:8080"

func main() {
	docker, err := createDockerClient()
	if err != nil {
		slog.Error("Error while connecting to docker daemon", "error", err)
		os.Exit(1)
	}

	runServer(docker)
}

func createDockerClient() (*client.Client, error) {
	docker, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate docker client: %v", err)
	}
	_, err = docker.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to docker daemon: %v", err)
	}
	return docker, nil
}

func runServer(docker *client.Client) {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.GET("/info", func(c *gin.Context) {
		// https://docs.docker.com/engine/api/v1.44/#tag/System/operation/SystemInfo
		info, err := docker.Info(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, info)
	})
	server.GET("/containers", func(c *gin.Context) {
		containers, err := docker.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var containersData []ContainerProperties
		for _, container := range containers {
			for i, name := range container.Names {
				container.Names[i] = strings.Replace(name, "/", "", -1)

			}
			realTime := time.Unix(container.Created, 0)
			realTimeString := realTime.Format(time.DateTime)
			containerData := ContainerProperties{
				Id:           container.ID,
				Names:        container.Names,
				Image:        container.Image,
				Status:       container.State,
				CreationDate: realTimeString,
			}
			containersData = append(containersData, containerData)
		}

		c.JSON(http.StatusOK, containersData)
		for _, ctr := range containers {
			fmt.Printf("Here \n")
			fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
		}
	})

	server.GET("/sse", handleSSE)
	server.POST("/send-data", handleSendData)

	slog.Info("Starting server", "listenAddr", listenAddr)
	server.Run(listenAddr)

}
func handleSSE(c *gin.Context) {
	containerID := c.Query("containerId")
	if containerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "containerId query parameter is required"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	log.Println("Client connected for container ID:", containerID)
	eventChan := make(chan string)
	clientsMu.Lock()
	clients[eventChan] = struct{}{}
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, eventChan)
		clientsMu.Unlock()
		close(eventChan)
	}()

	notify := c.Writer.CloseNotify()
	go func() {
		<-notify
		log.Println("Client disconnected")
		clientsMu.Lock()
		delete(clients, eventChan)
		clientsMu.Unlock()
	}()

	go streamContainerStats(containerID, eventChan)

	for {
		select {
		case data := <-eventChan:
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
		case <-notify:
			return
		}
	}
}

func handleSendData(c *gin.Context) {
	// This endpoint may be used to trigger the stats collection if needed
	c.JSON(http.StatusOK, gin.H{"message": "Endpoint not used directly"})
}

func streamContainerStats(containerID string, eventChan chan string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Printf("Failed to create Docker client: %v", err)
		return
	}
	stats, err := cli.ContainerStats(context.Background(), containerID, true)
	if err != nil {
		log.Printf("Failed to get container stats: %v", err)
		return
	}
	defer stats.Body.Close()

	decoder := json.NewDecoder(stats.Body)
	log.Println("Streaming decoder initialized for container ID:", containerID)
	for {
		var stat types.StatsJSON
		if err := decoder.Decode(&stat); err != nil {
			if err == io.EOF {
				log.Println("EOF reached, stopping stream for container ID:", containerID)
				break
			}
			log.Printf("Error decoding stats: %v", err)
			break
		}

		memoryUsage := stat.MemoryStats.Usage
		memoryLimit := stat.MemoryStats.Limit
		memoryPercent := float64(memoryUsage) / float64(memoryLimit) * 100.0

		containerStats := ContainerStats{
			CpuUsagePercent:    calculateCPUPercentUnix(stat),
			MemoryUsagePercent: memoryPercent,
			RxNetworkBytes:     stat.Networks["eth0"].RxBytes,
			TxNetworkBytes:     stat.Networks["eth0"].TxBytes,
		}

		statsJSON, err := json.Marshal(containerStats)
		if err != nil {
			log.Printf("Failed to encode stats to JSON: %v", err)
			continue
		}

		log.Printf("Broadcasting stats: %s", statsJSON)
		eventChan <- string(statsJSON)
		time.Sleep(500 * time.Microsecond) //better user experience in my machine
	}
}

type ContainerStats struct {
	CpuUsagePercent    float64 `json:"cpu_usage_percent"`
	MemoryUsagePercent float64 `json:"memory_usage_percent"`
	RxNetworkBytes     uint64  `json:"rx_network_bytes"`
	TxNetworkBytes     uint64  `json:"tx_network_bytes"`
}

var (
	clients   = make(map[chan string]struct{})
	clientsMu sync.Mutex // Mutex to synchronize access to clients map
)

func calculateCPUPercentUnix(stat types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		cpuDelta   = float64(stat.CPUStats.CPUUsage.TotalUsage) - float64(stat.PreCPUStats.CPUUsage.TotalUsage)

		systemDelta = float64(stat.CPUStats.SystemUsage) - float64(stat.PreCPUStats.SystemUsage)
	)
	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(stat.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent
}
