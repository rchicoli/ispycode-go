package main

import (
	"fmt"
)

// count from one to nine and skip five
func main() {
	number := 0
	for {
		number++
		if number > 9 {
			break
		}
		if number == 5 {
			continue
		}
		fmt.Print(number)
	}
}
