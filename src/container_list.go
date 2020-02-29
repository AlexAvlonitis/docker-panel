package main

import (
	"encoding/json"
)

const containerListPath = "http://localhost/containers/json"

func (d *DockerEngine) containerList() []Container {
	resp, err := d.client.Get(containerListPath)
	dieIf(err)

	var containers []Container
	err = json.NewDecoder(resp.Body).Decode(&containers)
	dieIf(err)

	return containers
}
