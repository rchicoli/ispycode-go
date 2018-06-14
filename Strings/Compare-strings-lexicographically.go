package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "A cat is a cat"
	str2 := "A cat is a dog"
	str3 := "A cat is a bat"

	// str1 is the same as str1
	x := strings.Compare(str1, str1)
	fmt.Println(x)

	// str2 should be less then str1
	x = strings.Compare(str1, str2)
	fmt.Println(x)

	// str3 should be greater then str1
	x = strings.Compare(str1, str3)
	fmt.Println(x)

}
