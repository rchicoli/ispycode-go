package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	mypath := "/usr/local/go/bin:/usr/local/bin:/usr/bin:/usr/sbin"
	paths := filepath.SplitList(mypath)
	for i := range paths {
		fmt.Println(paths[i])
	}
}
