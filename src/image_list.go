package main

import (
	"encoding/json"
)

const imageListPath = "http://localhost/images/json"

func (d *DockerEngine) imageList() []ImageSummary {
	resp, err := d.client.Get(imageListPath)
	dieIf(err)

	var images []ImageSummary
	err = json.NewDecoder(resp.Body).Decode(&images)
	dieIf(err)

	return images
}
