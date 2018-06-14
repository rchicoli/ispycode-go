package main

import (
	"fmt"
)

func main() {

	// Declare an integer array of five elements.
	// Initialize each element with a specific value.
	array := [5]int{0, 10, 20, 30, 40}

	// Change the value at index 2.
	array[2] = 22

	// print contents
	fmt.Println(array)

	// print content of index 2
	fmt.Println(array[2])

}
