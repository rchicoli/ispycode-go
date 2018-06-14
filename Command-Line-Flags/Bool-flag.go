package main

import (
	"flag"
	"fmt"
)

func main() {

	b := flag.Bool("mybool", true, "some boolean")

	// Parse the flags.
	flag.Parse()

	// Get bool from the bool pointer.
	value := *b
	fmt.Println("mybool:", value)
}
