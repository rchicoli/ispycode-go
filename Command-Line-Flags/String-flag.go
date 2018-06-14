package main

import (
	"flag"
	"fmt"
)

func main() {

	mystring := flag.String("mystring", "hello world", "a short string")

	// Parse the flags.
	flag.Parse()

	// Get string from the pointer.
	value := *mystring
	fmt.Println("mystring:", value)
}
