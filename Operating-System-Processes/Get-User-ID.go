package main

import (
	"fmt"
	"os"
)

func main() {
	uid := os.Getuid()
	fmt.Println("uid:", uid)
}
