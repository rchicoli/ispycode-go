package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	targpath := "/home/dennis/projects"
	basepath := "/home/dennis/projects/golang/file1"

	relpath, _ := filepath.Rel(targpath, basepath)
	fmt.Println("Relative Path:", relpath)
}
