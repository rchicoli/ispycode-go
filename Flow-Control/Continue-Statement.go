package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			// skip even numbers
			continue
		}
		fmt.Println(i)
	}
}
