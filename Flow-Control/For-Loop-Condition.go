package main

import "fmt"

func main() {

	valid := true
	i := 1

	for valid {

		if i <= 5 {
			valid = false
		}

		fmt.Println(i)
		i++
	}
}
