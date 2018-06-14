package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "The cat in the hat"

	b1 := strings.Contains(str, "cat")
	fmt.Println(b1)

	b2 := strings.Contains(str, "dog")
	fmt.Println(b2)
}
