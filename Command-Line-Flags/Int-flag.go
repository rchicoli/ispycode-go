package main

import (
	"flag"
	"fmt"
)

func main() {

	number := flag.Int("number", 5, "cool number")

	// Parse the flags.
	flag.Parse()

	// Get int from the Int pointer.
	value := *number
	fmt.Println("number:", value)
}
