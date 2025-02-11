package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) serverSendEventsHandler(w http.ResponseWriter, r *http.Request) {
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
