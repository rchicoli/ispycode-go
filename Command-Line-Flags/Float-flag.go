package main

import (
	"flag"
	"fmt"
)

func main() {

	myfloat := flag.Float64("myfloat", 1.234567, "a float")

	// Parse the flags.
	flag.Parse()

	// Get float from the float pointer.
	value := *myfloat
	fmt.Println("myfloat:", value)
}
