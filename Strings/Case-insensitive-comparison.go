package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "cat"
	str2 := "Cat"

	// returns true because cases are ignored.
	b := strings.EqualFold(str1, str2)

	fmt.Println(b)
}
