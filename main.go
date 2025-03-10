package main

import (
	"log"

	"github.com/XaviFelix/log-processor.git/fileio"
)

func main() {

	err := fileio.ReadFile("super.txt")
	if err != nil {
		// printing error to console and stoppping execution
		log.Fatalf("error from main: %v", err)
	}

}
