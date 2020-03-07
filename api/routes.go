package main

import "github.com/gorilla/mux"

// NewRouter initialize the router with its routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	initRoutes(router)
	return router
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/images", imageList).Methods("GET")
	router.HandleFunc("/containers", containerList).Methods("GET")
	router.HandleFunc("/networks", networkList).Methods("GET")
	router.HandleFunc("/volumes", volumeList).Methods("GET")
}
