package main

import (
	"fmt"
)

func main() {

	// Declare an integer array of six elements.
	// Initialize index 2 and 4 with specific values.
	// The rest of the elements contain their zero value.
	array := [6]int{2: 20, 4: 40}

	// print contents
	fmt.Println(array)

	// print size (length)
	fmt.Println(len(array))
}
