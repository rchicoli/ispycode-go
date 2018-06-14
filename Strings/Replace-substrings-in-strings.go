package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "A cat is a cat is a cat is a cat"
	fmt.Println(str1)

	// replace first occurrence
	str2 := strings.Replace(str1, "cat", "dog", 1)
	fmt.Println(str2)

	// replace frst two occurrences
	str3 := strings.Replace(str1, "cat", "dog", 2)
	fmt.Println(str3)

	// replace all occurrences
	str4 := strings.Replace(str1, "cat", "dog", -1)
	fmt.Println(str4)
}
