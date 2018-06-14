package main

import "fmt"

func main() {

	var z []int
	fmt.Println(z)

	fmt.Println("Length:", len(z))
	fmt.Println("Capacity:", cap(z))

	if z == nil {
		fmt.Println("It is nil")
	}
}
