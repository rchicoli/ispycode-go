package main

import "fmt"

func main() {

	f := 5.01234567890123456789

	// default width, default precision
	fmt.Printf("Float 1: %f \n", f)

	// width 9, default precision
	fmt.Printf("Float 2: %9f \n", f)

	// default width, precision 2
	fmt.Printf("Float 3: %.2f \n", f)

	// width 9, precision 2
	fmt.Printf("Float 4: %9.2f \n", f)

	// width 9, precision 0
	fmt.Printf("Float 5: %9.f \n", f)
}
