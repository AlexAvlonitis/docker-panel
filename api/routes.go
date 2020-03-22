package main

import "github.com/gorilla/mux"

// NewRouter initialize the router with its routes
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	initRoutes(r)
	return r
}

func initRoutes(r *mux.Router) {
	r.HandleFunc("/images", ImageList)
	r.HandleFunc("/images/{imageID}", ImageInspect)
	r.HandleFunc("/containers", ContainerList)
	r.HandleFunc("/containers/{containerID}", ContainerInspect)
	r.HandleFunc("/networks", NetworkList)
	r.HandleFunc("/networks/{networkID}", NetworkInspect)
	r.HandleFunc("/volumes", VolumeList)
	r.HandleFunc("/volumes/{volumeID}", VolumeInspect)
}
