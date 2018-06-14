package main

import "fmt"

func main() {

	i := 123456

	fmt.Printf("The number in octal is %o \n", i)

	// alternate format: add leading 0 for octal (%#o)
	fmt.Printf("The number in octal is %#o \n", i)
}
