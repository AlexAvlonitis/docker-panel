package main

import (
	"fmt"
)

func main() {
	de := newDockerEngine()

	fmt.Println(de.imageList())
	fmt.Println(de.containerList())
}
