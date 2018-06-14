package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Thing struct {
	Name string
	Size string
}

func main() {

	csvfile, err := os.Open("/tmp/myfile.csv")
	checkError(err)
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var thing Thing
	var things []Thing

	for _, record := range rawCSVdata {
		thing.Name = record[0]
		thing.Size = record[1]
		things = append(things, thing)
	}

	fmt.Println(things)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
