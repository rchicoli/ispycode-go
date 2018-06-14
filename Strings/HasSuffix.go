package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "bingo"
	str2 := "bongo"
	str3 := "gono"

	suffix := "go"

	b := strings.HasSuffix(str1, suffix)
	fmt.Println(str1, ":", b)

	b = strings.HasSuffix(str2, suffix)
	fmt.Println(str2, ":", b)

	b = strings.HasSuffix(str3, suffix)
	fmt.Println(str3, ":", b)
}
