package main

import (
	"fmt"
)

func main() {

	var array1 [5]int

	array2 := [5]int{0, 1, 2, 3, 4}

	// Copy the values from array2 into array1.
	array1 = array2

	// change cell in array1
	array1[2] = 22

	// print contents
	fmt.Println(array1)
	fmt.Println(array2)

}
