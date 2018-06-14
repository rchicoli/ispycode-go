package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "the cat in the hat"
	fmt.Println(str)

	title := strings.Title(str)
	fmt.Println(title)
}
