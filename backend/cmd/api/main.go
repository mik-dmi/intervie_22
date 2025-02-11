package main

import (
	"context"
	"encoding/json"
	"kinexon/containerruntime/utils"
	"sync"

	"io"
	"log"
	"net/http"

	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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
		log.Fatal("Error while connecting to docker daemon", err)
	}

	config := config{
		addr:       ":8080",
		dockerAddr: ":8080",
	}

	app := &application{
		config:       config,
		dockerClient: docker,
	}

	mux := app.mount()
	log.Fatal(app.runServer(mux))

}

func (app *application) sendDataHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
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
			CpuUsagePercent:    utils.CalculateCPUPercentUnix(stat),
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
