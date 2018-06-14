package main

import "fmt"

func main() {

	max := int32(0)
	for i := int32(0); i >= 0; i++ {
		max = i
	}
	fmt.Println("Max:", max)

	min := int32(0)
	for i := int32(0); i <= 0; i-- {
		min = i
	}
	fmt.Println("Min:", min)
}
