package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	abs := filepath.IsAbs("../../Playground/hello.go")
	fmt.Println("Is absolute:", abs)
}
