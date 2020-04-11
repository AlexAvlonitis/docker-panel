package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gorilla/mux"
)

// InfoHandler returns `docker info` as json
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	info, err := cli.Info(context.Background())
	dieIfErr(err)

	renderJSON(w, info)
}

// ImageList returns `docker image ls` as json
func ImageList(w http.ResponseWriter, r *http.Request) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	dieIfErr(err)

	renderJSON(w, images)
}

// ImageInspect returns `docker image inspect {id}` as json
func ImageInspect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	images, _, err := cli.ImageInspectWithRaw(context.Background(), params["imageID"])
	dieIfErr(err)

	renderJSON(w, images)
}

// ContainerList returns `docker container ls` as json
func ContainerList(w http.ResponseWriter, r *http.Request) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	dieIfErr(err)

	renderJSON(w, containers)
}

// ContainerInspect returns `docker volume inspect {id}` as json
func ContainerInspect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	containers, _, err := cli.ContainerInspectWithRaw(
		context.Background(),
		params["containerID"],
		true,
	)
	dieIfErr(err)

	renderJSON(w, containers)
}

// NetworkList returns `docker network ls` as json
func NetworkList(w http.ResponseWriter, r *http.Request) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	dieIfErr(err)

	renderJSON(w, networks)
}

// NetworkInspect returns `docker network inspect {id}` as json
func NetworkInspect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	insOpts := types.NetworkInspectOptions{}
	networks, _, err := cli.NetworkInspectWithRaw(
		context.Background(), params["networkID"], insOpts,
	)
	dieIfErr(err)

	renderJSON(w, networks)
}

// VolumeList returns `docker volume ls` as json
func VolumeList(w http.ResponseWriter, r *http.Request) {
	a := filters.Args{}
	volumes, err := cli.VolumeList(context.Background(), a)
	dieIfErr(err)

	renderJSON(w, volumes)
}

// VolumeInspect returns `docker volume inspect {id}` as json
func VolumeInspect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	volumes, _, err := cli.VolumeInspectWithRaw(context.Background(), params["volumeID"])
	dieIfErr(err)

	renderJSON(w, volumes)
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	initHeaders(&w)
	json.NewEncoder(w).Encode(data)
}

func initHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}
