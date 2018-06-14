package main

import "fmt"

func main() {

	var a = 0x8
	fmt.Printf("%b \n", a)

	var b = 0xa
	fmt.Printf("%b \n", b)

	c := a | b
	fmt.Printf("%b \n", c)
}
