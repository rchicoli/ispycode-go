package main

import (
	"io"
	"log"
	"os"
)

func main() {

	file, e := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}

	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)

	log.Println("message 1")

}
