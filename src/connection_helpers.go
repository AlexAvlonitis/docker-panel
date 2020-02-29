package main

import (
	"fmt"
	"os"
)

func dieIf(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
}
