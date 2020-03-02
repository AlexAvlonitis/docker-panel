package main

import "github.com/gorilla/mux"

// InitRouter initialize the router with its routes
func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	initRoutes(router)
	return router
}

func initRoutes(router *mux.Router) {
	router.HandleFunc("/images", imageList).Methods("GET")
	router.HandleFunc("/containers", containerList).Methods("GET")
}
