package main

import (
	"fmt"
	"sort"
)

func main() {

	colors := map[string]int{
		"red":    5,
		"yellow": 2,
		"blue":   3,
	}
	fmt.Println(colors)

	// get the list of keys and sort them
	keys := []string{}
	for key := range colors {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, val := range keys {
		fmt.Println(val, colors[val])
	}
}
