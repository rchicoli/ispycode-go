package main

import "fmt"

func main() {

	var i = 0x8
	fmt.Printf("%b \n", i)

	// shift right two
	i >>= 2
	fmt.Printf("%b \n", i)
}
