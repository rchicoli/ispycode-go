package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
		"green":  3,
	}
	fmt.Println(colors)

	delete(colors, "orange")
	fmt.Println(colors)
}
