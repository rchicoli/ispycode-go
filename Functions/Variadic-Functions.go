package main

import "fmt"

func main() {
	fmt.Println("Total:", sum(1))
	fmt.Println("Total:", sum(1, 2))
	fmt.Println("Total:", sum(1, 2, 3))
	fmt.Println("Total:", sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
