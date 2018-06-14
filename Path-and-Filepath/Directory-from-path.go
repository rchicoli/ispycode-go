package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir := filepath.Dir("/some-path/to-file/hello.go")
	fmt.Println("Directory:", dir)
}
