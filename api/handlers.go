package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
)

func initHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	initHeaders(&w)
	json.NewEncoder(w).Encode(data)
}

func imageList(w http.ResponseWriter, r *http.Request) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	dieIfErr(err)

	renderJSON(w, images)
}

func containerList(w http.ResponseWriter, r *http.Request) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	dieIfErr(err)

	renderJSON(w, containers)
}

func networkList(w http.ResponseWriter, r *http.Request) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	dieIfErr(err)

	renderJSON(w, networks)
}
