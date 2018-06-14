package main

import (
	"fmt"
	"os"
)

func main() {

	// The first element os.Args[0] is the name of the command
	command := os.Args[0]
	fmt.Println("Command:", command)

	// the rest are arguments
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Printf("Arg %d: %s \n", i, args[i])
	}
}
