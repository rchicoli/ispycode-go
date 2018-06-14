package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("/tmp")

	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println(mydir)
	}
}
