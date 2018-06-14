package main

import "fmt"

func main() {
	a := []int{333, 444, 555}

	// index and value
	for i, v := range a {
		fmt.Println("Index:", i, "Value:", v)
	}

	// index only
	for i, _ := range a {
		fmt.Println("Index:", i)
	}

	// value only
	for _, v := range a {
		fmt.Println("Value:", v)
	}
}
