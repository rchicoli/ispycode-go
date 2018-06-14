package main

import "fmt"

func main() {

	colors := map[string]int{
		"red":    1,
		"orange": 2,
		"yellow": 3,
	}

	for key, value := range colors {
		fmt.Println(key, ":", value)
	}
}
