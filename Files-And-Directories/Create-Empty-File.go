package main

import (
	"fmt"
	"os"
)

func main() {
	newFile, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	newFile.Close()
}
