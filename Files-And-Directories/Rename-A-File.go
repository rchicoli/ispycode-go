package main

import (
	"log"
	"os"
)

func main() {
	oldFile := "file1.txt"
	newFile := "file2.txt"
	err := os.Rename(oldFile, newFile)
	if err != nil {
		log.Fatal(err)
	}
}
