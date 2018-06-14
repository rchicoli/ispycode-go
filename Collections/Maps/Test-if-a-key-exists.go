package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
		"green":  3,
	}

	value, ok := colors["red"]
	if ok == true {
		fmt.Println(value)
	}

	value, ok = colors["yellow"]
	if ok != true {
		fmt.Println("Yellow does not exist")
	}
}
