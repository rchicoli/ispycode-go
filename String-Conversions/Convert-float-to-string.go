package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 1.23456789

	// float
	// 5 bit precision
	// 64 bit
	str := strconv.FormatFloat(f, 'f', 5, 64)

	fmt.Println(str)
}
