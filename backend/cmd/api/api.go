package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/docker/client"
)

type application struct {
	config       config
	dockerClient *client.Client
}

type config struct {
	addr       string
	dockerAddr string
}

type dockerClient struct {
	addr string
}

func createDockerClient() (*client.Client, error) {
	docker, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate docker client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err = docker.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to docker daemon: %v", err)
	}
	return docker, nil
}

func (app *application) mount() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /info", app.dockerInfoHandler)
	router.HandleFunc("GET /containers", app.containersHandler)
	router.HandleFunc("GET /sse", app.serverSendEventsHandler)
	router.HandleFunc("GET /send-data", app.sendDataHandler)

	return router

}

func (app *application) runServer(mux *http.ServeMux) error {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server is running at", app.config.addr)
	return s.ListenAndServe()

}
