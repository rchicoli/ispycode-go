package main

import (
	"log"
	"os"
)

func main() {

	err := os.Chmod("myfile.txt", 0777)
	if err != nil {
		log.Println(err)
	}
}
