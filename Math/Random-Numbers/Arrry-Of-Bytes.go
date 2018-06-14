package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var randomBytes = make([]byte, 10)
	length, err := rand.Read(randomBytes)

	if err == nil {
		fmt.Println("Length:", length)
		fmt.Println(randomBytes)
	}
}
