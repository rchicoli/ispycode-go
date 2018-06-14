package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "The cat in the hat"
	i := strings.Index(str, "cat")
	fmt.Println(i)
}
