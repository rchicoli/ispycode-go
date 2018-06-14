package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
		"yellow": 3,
	}

	// Retrieve the value for the key "red".
	value, exists := colors["red"]
	if exists {
		fmt.Println("red:", value)
	} else {
		fmt.Println("red: does not exists")
	}

	// Retrieve the value for the key "green".
	value, exists = colors["green"]
	if exists {
		fmt.Println("green:", value)
	} else {
		fmt.Println("green: does not exists")
	}
}
