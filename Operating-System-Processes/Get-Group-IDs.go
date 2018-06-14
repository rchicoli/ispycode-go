package main

import (
	"fmt"
	"os"
)

func main() {
	groups, err := os.Getgroups()
	if err == nil {
		fmt.Println("groups:", groups)
	}
}
