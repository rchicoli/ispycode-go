package main

import "fmt"

func main() {

	height := 10

	// top line
	for y := 1; y < height; y++ {
		fmt.Print(" ")
	}
	fmt.Println("*")

	// middle lines
	for x := 2; x < height; x++ {
		for y := 1; y <= height-x; y++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for y := 1; y < ((x * 2) - 2); y++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}

	// bottom line
	for y := 1; y < (height*2)-1; y++ {
		fmt.Print("*")
	}
	fmt.Println("*")
}
