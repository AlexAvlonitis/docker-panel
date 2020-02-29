package main

import (
	"net"
	"net/http"
)

// DockerUnixSocket Docker socket connection
const DockerUnixSocket = "/var/run/docker.sock"

// DockerEngine main client for the docker engine
type DockerEngine struct {
	client *http.Client
}

func unixDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", "/var/run/docker.sock")
}

func newDockerEngine() DockerEngine {
	tr := &http.Transport{
		Dial: unixDial,
	}
	c := &http.Client{Transport: tr}

	return DockerEngine{client: c}
}
