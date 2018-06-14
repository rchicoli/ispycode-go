package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abs, err := filepath.Abs("./hello.go")
	if err == nil {
		fmt.Println("Absolute:", abs)
	}
}
