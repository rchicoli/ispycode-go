package main

import (
	"fmt"
	"os"
)

func main() {

	// get FileInfo structure describing file
	fileInfo, err := os.Stat("myfile.txt")

	// check if there is an error
	if err != nil {

		// check if error is file does not exist
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		}

	} else {
		fmt.Println(fileInfo.Name(), "exists")
	}
}
