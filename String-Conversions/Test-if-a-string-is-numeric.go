package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(IsNumeric("1"))
	fmt.Println(IsNumeric("12.345"))
	fmt.Println(IsNumeric("NOT"))
}

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
