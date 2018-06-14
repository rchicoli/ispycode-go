package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkfunc(path string, info os.FileInfo, err error) error {
	fmt.Println("Path:", path)
	return nil
}

func main() {
	filepath.Walk("/tmp/mydir", walkfunc)
}
