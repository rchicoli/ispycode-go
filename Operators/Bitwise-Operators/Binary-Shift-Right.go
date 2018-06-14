package main

import "fmt"

func main() {

	var a = 0x10
	fmt.Printf("%05b \n", a)

	b := a >> 3
	fmt.Printf("%05b \n", b)
}
