package main

import (
	"fmt"
)

func main() {

	// Declare a string pointer array of three elements.
	var array1 [3]*string

	// Declare a second string pointer array of three elements.
	// Initialize the array with string pointers.
	array2 := [3]*string{new(string), new(string), new(string)}

	// Add colors to each element
	*array2[0] = "Red"
	*array2[1] = "Orange"
	*array2[2] = "Yellow"

	// Copy the values from array2 into array1.
	// After the copy, you have two arrays pointing to the same strings,
	array1 = array2

	// Print contents
	fmt.Println(*array1[0], *array1[1], *array1[2])
	fmt.Println(*array2[0], *array2[1], *array2[2])

	// Change the value of one of the elements
	*array1[1] = "Blue"

	// Print contents
	fmt.Println(*array1[0], *array1[1], *array1[2])
	fmt.Println(*array2[0], *array2[1], *array2[2])

}
