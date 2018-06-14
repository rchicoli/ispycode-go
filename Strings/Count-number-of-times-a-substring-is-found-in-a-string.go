package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "cat dog cat dog cat"
	str2 := "cat"

	fmt.Printf("%v \n", strings.Count(str1, str2))
}
