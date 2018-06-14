package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "New York"
	str2 := "New Mexico"
	str3 := "South Carolina"

	prefix := "New"

	b := strings.HasPrefix(str1, prefix)
	fmt.Println(str1, ":", b)

	b = strings.HasPrefix(str2, prefix)
	fmt.Println(str2, ":", b)

	b = strings.HasPrefix(str3, prefix)
	fmt.Println(str3, ":", b)
}
