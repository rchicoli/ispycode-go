package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var data = [][]string{{"aaa", "bbb", "ccc"}, {"ddd", "eee", "fff"}}

func main() {
	file, err := os.Create("/tmp/result.csv")
	checkError(err)
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, value := range data {
		err := writer.Write(value)
		checkError(err)
	}

	defer writer.Flush()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
