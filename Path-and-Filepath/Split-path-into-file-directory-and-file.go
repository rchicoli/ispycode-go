package main

import (
	"fmt"
	"path"
)

func main() {

	dir, file := path.Split("/Users/dennis/ISPY/GO/example.go")
	fmt.Println("Dir:", dir)
	fmt.Println("File:", file)
}
