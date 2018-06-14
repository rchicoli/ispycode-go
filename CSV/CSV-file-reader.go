package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/tmp/sample.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// NewReader returns a new Reader that reads from file.
	reader := csv.NewReader(file)

	lineCount := 0
	for {
		// Read reads one record from the reader. The record is a slice of strings
		// with each string representing one field.
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		for i := 0; i < len(record); i++ {
			fmt.Println(" ", record[i])
		}
		fmt.Println()
		lineCount += 1
	}
}
