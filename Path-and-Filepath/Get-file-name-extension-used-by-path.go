package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	extension := filepath.Ext("/tmp/hello.go")
	fmt.Println("Extension 1:", extension)

	extension = filepath.Ext("./hello.jpg")
	fmt.Println("Extension 2:", extension)
}
