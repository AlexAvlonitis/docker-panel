package main

import (
	"fmt"
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
	r := NewRouter()

	fmt.Println("Server running at " + httpAddr)
	dieIfErr(http.ListenAndServe(httpAddr, r))
}
