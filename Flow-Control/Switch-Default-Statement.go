package main

import "fmt"

func main() {
	x := 5
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Unknown")
	}
}
