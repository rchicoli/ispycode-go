package main

import "fmt"

func main() {

	// set a variable
	i := 100
	fmt.Println(i)

	// get the address of i
	p := &i
	fmt.Println(p)

	// print the contents of the address
	fmt.Println(*p)

}
