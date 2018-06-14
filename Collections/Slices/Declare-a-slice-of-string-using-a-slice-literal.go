package main

import (
	"fmt"
)

func main() {

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	slice := []string{"Red", "Orange", "Yellow", "Green", "Blue"}

	// Print contents, length and capacity
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
