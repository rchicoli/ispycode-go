package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	base := filepath.Base("/home/dennis/IdeaProjects/Playground/hello.go")
	fmt.Println("Base:", base)
}
