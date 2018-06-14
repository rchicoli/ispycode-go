package main

import "fmt"

func main() {

	colors := make(map[string]int)
	fmt.Println(colors)

	colors["red"] = 1
	colors["orange"] = 2

	fmt.Println(colors)
}
