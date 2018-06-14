package main

import "fmt"

func main() {

	var slice []int

	fmt.Println(slice)
	fmt.Println("Capacity:", cap(slice))
	fmt.Println("Length:", len(slice))
}
