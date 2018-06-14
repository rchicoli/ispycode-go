package main

import "fmt"

func main() {

	colors := make(map[int]string)
	colors[1] = "Red"
	colors[2] = "Orange"
	colors[3] = "Yellow"

	fmt.Println(colors)
	fmt.Println(colors[1])
}
