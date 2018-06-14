package main

import "fmt"

func main() {

	count := 1

	// This loop continues infinitely until broken.
	for {

		// Break if count is greater the 10.
		if count > 10 {
			break
		}
		fmt.Println(count)
		count++
	}
}
