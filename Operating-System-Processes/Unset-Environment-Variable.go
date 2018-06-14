package main

import (
	"fmt"
	"os"
)

func main() {
	myenv := os.Getenv("MYENV")
	fmt.Println("myenv:", myenv)

	os.Unsetenv("MYENV")

	myenv = os.Getenv("MYENV")
	fmt.Println("myenv:", myenv)
}
