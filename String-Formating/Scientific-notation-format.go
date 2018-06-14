package main

import "fmt"

func main() {

	f := 1.2345678901234567890

	// lower case
	fmt.Printf("Scientific notation %e \n", f)

	// upper case
	fmt.Printf("Scientific notation %E \n", f)

	// decimal point but no exponent
	// %F is synonym for %f
	fmt.Printf("Scientific notation %f \n", f)
	fmt.Printf("Scientific notation %F \n", f)
}
