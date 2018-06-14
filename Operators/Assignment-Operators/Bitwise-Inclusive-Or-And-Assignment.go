package main

import "fmt"

func main() {

	var i = 0x8
	fmt.Printf("%b \n", i)

	i |= 0x2
	fmt.Printf("%b \n", i)
}
