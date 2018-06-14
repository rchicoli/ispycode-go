package main

import (
	"fmt"
)

func main() {

	// Declare a two dimensional integer array of three elements
	// by two elements.
	var array1 [3][2]int

	// Print contents
	fmt.Println(array1)

	// Use an array literal to declare and initialize a two
	// dimensional integer array.
	array2 := [3][2]int{{10, 11}, {20, 21}, {30, 31}}

	// Print contents
	fmt.Println(array2)

	// Declare and initialize index 0 and 1 of the outer array.
	array3 := [3][2]int{0: {10, 11}, 2: {30, 31}}

	// Print contents
	fmt.Println(array3)

	// Declare and initialize individual elements of the outer
	// and inner array.
	array4 := [3][2]int{0: {0: 10}, 1: {0: 20}, 2: {1: 31}}

	// Print contents
	fmt.Println(array4)

}
