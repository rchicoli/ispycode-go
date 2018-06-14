package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "the cat in the hat"

	// string into an array of words
	f := strings.Fields(str)

	// iterate over array
	for _, v := range f {
		fmt.Println(v)
	}
}
