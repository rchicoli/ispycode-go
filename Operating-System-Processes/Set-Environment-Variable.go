package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("MYENV", "abcdefg")

	myenv := os.Getenv("MYENV")
	fmt.Println("myenv:", myenv)
}
