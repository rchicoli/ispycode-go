package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "abc..."
	fmt.Println(str1)

	str2 := strings.Repeat(str1, 5)
	fmt.Println(str2)
}
