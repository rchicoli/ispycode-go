package main

import "fmt"

func floyd(rows int) {
	number := 1
	for a := 1; a <= rows; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%-2d ", number)
			number++
		}
		fmt.Println("")
	}
}

func main() {
	floyd(10)
}
