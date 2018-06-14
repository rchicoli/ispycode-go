package main

import (
	"fmt"
	"sort"
)

func main() {

	colors := map[string]int{
		"red":    50,
		"yellow": 16,
		"blue":   30,
	}
	fmt.Println(colors)

	// used to switch key and value
	hack := map[int]string{}
	hackkeys := []int{}
	for key, val := range colors {
		hack[val] = key
		hackkeys = append(hackkeys, val)
	}
	sort.Ints(hackkeys)

	for _, val := range hackkeys {
		fmt.Println(hack[val], val)
	}
}
