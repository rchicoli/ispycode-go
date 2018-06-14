package main

import (
	"fmt"
)

func main() {

	// Create a slice of a slice of integers.
	slice := [][]int{{10}, {100, 200}}

	// Print contents
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Append the value of 20 to the first slice of integers.
	slice[0] = append(slice[0], 20)

	// Print contents
	fmt.Println(slice)

	// Print contents first slice
	fmt.Println(slice[0])
	fmt.Println(len(slice[0]))
	fmt.Println(cap(slice[0]))

	// Print contents second slice
	fmt.Println(slice[1])
	fmt.Println(len(slice[1]))
	fmt.Println(cap(slice[1]))

}
