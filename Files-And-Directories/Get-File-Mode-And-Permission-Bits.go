package main

import (
	"fmt"
	"os"
)

func main() {
	file := "myfile.txt"
	fileInfo, _ := os.Stat(file)
	mode := fileInfo.Mode()

	fmt.Println(file, "mode is ", mode)
}
