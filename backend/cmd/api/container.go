package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/gin-gonic/gin"
)

func (app *application) containersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := app.dockerClient.ContainerList(context.Background(), container.ListOptions{})
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
}

func (app *application) dockerInfoHandler(w http.ResponseWriter, r *http.Request) {

	info, err := app.dockerClient.Info(context.Background())
	if err != nil {
		http.Error(w, "Failed to retrieve Docker info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
