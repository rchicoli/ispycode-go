package main

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Close file
	file.Close()

	// Create .gz file to write to
	outputFile, err := os.Create("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	// Create a gzip writer on top of file writer
	gzipWriter := gzip.NewWriter(outputFile)

	// Write out compressed data
	_, err = gzipWriter.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Close writer
	gzipWriter.Close()

	log.Println("Compressed data written to file.")
}
