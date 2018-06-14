package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"the", "cat", "in", "the", "hat"}
	fmt.Println(str1)

	str2 := strings.Join(str1, " ")
	fmt.Println(str2)

	str3 := strings.Join(str1, ", ")
	fmt.Println(str3)
}
