package main

import "fmt"

func main() {
	i := 123456

	// lowercase
	fmt.Printf("The number in hex is %x \n", i)

	// uppercase
	fmt.Printf("The number in hex is %X \n", i)

	// alternate format: add leading 0x for hex
	fmt.Printf("The number in hex is %#x \n", i)
	fmt.Printf("The number in hex is %#X \n", i)
}
