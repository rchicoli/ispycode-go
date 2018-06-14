package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	mypath := "C:/examples/golang/example.go"
	fmt.Println("Volume:", filepath.VolumeName(mypath))
}
