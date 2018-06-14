package main

import "fmt"

func main() {

	i1 := 10
	i2 := -20

	// always print a sign for numeric values
	fmt.Printf("Signed number: %+d \n", i1)
	fmt.Printf("Signed number: %+d \n", i2)

	// leave a space for elided sign in numbers
	fmt.Printf("Signed number: % d \n", i1)
	fmt.Printf("Signed number: % d \n", i2)

}
