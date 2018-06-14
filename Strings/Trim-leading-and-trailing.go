package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "  The cat in the hat   "
	fmt.Println(str1)

	// trim left
	str2 := strings.TrimLeft(str1, " ")
	fmt.Println(str2)

	// trim right
	str3 := strings.TrimRight(str1, " ")
	fmt.Println(str3)

	// trim both sides
	str4 := strings.Trim(str1, " ")
	fmt.Println(str4)

	// using TrimSpace
	str5 := strings.TrimSpace(str1)
	fmt.Println(str5)

}
