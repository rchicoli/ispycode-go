package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "A cat is a cat is a cat."
	i := strings.LastIndex(str, "cat")
	fmt.Println(i)
}
