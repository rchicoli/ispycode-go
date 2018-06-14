package main

import (
	"fmt"
	"strings"
)

// Parse a CSV string into an array

func main() {
	str := "aaa,bbb,ccc,ddd"
	arr := strings.Split(str, ",")
	fmt.Println(arr)
}
