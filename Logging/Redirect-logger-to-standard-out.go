package main

import (
	"log"
	"os"
)

func main() {
	log.Println("message 1")

	// redirect logger to standard out
	log.SetOutput(os.Stdout)

	log.Println("message 2")
}
