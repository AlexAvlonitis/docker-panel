package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/client"
)

const httpAddr = ":1234"

var cli, cliErr = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func dieIfErr(err error) {
	if err != nil {
		panic(cliErr)
	}
}

func main() {
	dieIfErr(cliErr)
	router := InitRouter()

	fmt.Println("Server running at " + httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, router))
}
