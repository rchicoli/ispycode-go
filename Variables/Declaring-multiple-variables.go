package main

import (
	"fmt"
)

func main() {

	var a, b int = 1, 2
	fmt.Println(a, b)

	// shorthand for declaring and initializing a variable
	c, d := 10, 11
	fmt.Println(c, d)

	var str1, str2 = "Hello", "World"
	fmt.Println(str1)
	fmt.Println(str2)
}
