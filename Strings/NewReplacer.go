package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "A cat and a dog are pets"
	fmt.Println(str)

	// Create replacer with pairs as arguments.
	r := strings.NewReplacer(
		"cat", "lion",
		"dog", "wolf",
		"pets", "wild")

	str = r.Replace(str)
	fmt.Println(str)

}
