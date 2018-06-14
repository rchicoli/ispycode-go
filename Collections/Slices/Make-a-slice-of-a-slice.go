package main

import "fmt"

func main() {

	// Create a slice of integers.
	// Contains a length and capacity of 5 elements.
	slice := []int{10, 20, 30, 40, 50}

	fmt.Println(slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Create a new slice.
	// Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]

	fmt.Println(newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))
}
