package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
		"yellow": 3,
	}

	fmt.Println(colors)

	// change orange
	colors["orange"] = 99
	fmt.Println(colors)
}
