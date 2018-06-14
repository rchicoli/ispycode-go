package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
}
