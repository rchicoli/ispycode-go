package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
	}

	fmt.Println(colors)

	// add element
	colors["yellow"] = 3
	fmt.Println(colors)
}
