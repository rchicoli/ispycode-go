package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	clean := filepath.Clean("/home/dennis/../dennis/IdeaProjects/Playground/hello.go")
	fmt.Println("Clean:", clean)
}
