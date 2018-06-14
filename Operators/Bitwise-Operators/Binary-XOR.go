package main

import "fmt"

func main() {

	var a = 0x8
	fmt.Printf("%04b \n", a)

	var b = 0x2
	fmt.Printf("%04b \n", b)

	c := a ^ b
	fmt.Printf("%04b \n", c)
}
