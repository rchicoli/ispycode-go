package main

import "fmt"

func main() {
	defer cleanup()
	fmt.Println("hello")
}

func cleanup() {
	fmt.Println("world")
}
