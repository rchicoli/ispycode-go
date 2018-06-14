package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file1, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	file2, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()

	count, err := io.Copy(file2, file1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Copied %d bytes.\n", count)

	err = file2.Sync()
	if err != nil {
		fmt.Println(err)
	}
}
