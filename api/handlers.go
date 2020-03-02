package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func imageList(w http.ResponseWriter, r *http.Request) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	dieIfErr(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func containerList(w http.ResponseWriter, r *http.Request) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	dieIfErr(err)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}
