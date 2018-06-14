package main

import (
	"log"
	"os"
)

func main() {

	file, e := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if e != nil {
		log.Fatalln("Failed to open log file")
	}

	log.SetOutput(file)

	log.Println("message 1")

}
