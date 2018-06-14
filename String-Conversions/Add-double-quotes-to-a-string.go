package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "hello world"
	quote1 := strconv.Quote(str1)
	fmt.Println(str1)
	fmt.Println(quote1)

	str2 := "hello \t world"
	quote2 := strconv.Quote(str2)
	fmt.Println(str2)
	fmt.Println(quote2)
}
