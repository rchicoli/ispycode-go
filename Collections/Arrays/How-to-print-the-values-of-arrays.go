package main

import (
	"fmt"
)

func main() {

	array := []int{10, 20, 30, 40, 50, 60}

	// using Println
	fmt.Println(array)

	// using Printf
	fmt.Printf("%v \n", array)

	// iterate index and value
	for i, v := range array {
		fmt.Println("Index:", i, "Value:", v)
	}

}
