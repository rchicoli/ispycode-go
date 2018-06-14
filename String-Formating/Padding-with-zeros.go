package main

import "fmt"

func main() {

	i1 := 10
	i2 := 6666
	f1 := 123.34

	fmt.Printf("Padding: %08d left\n", i1)
	fmt.Printf("Padding: %08d left\n", i2)
	fmt.Printf("Padding: %08.2f left\n", f1)
}
