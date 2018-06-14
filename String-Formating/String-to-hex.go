package main

import "fmt"

func main() {

	str := "Hello"

	// base 16, lower-case, two characters per byte
	fmt.Printf("String to hex %x \n", str)

	// base 16, upper-case, two characters per byte
	fmt.Printf("String to hex %X \n", str)
}
