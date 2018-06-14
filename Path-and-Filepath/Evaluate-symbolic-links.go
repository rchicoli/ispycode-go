package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path, _ := filepath.EvalSymlinks("/tmp/myfile2")
	fmt.Println("Path", path)
}
