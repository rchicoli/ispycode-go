package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Capital A
	a := strconv.IsPrint(65)
	fmt.Println(a)

	// Bell
	bel := strconv.IsPrint(7)
	fmt.Println(bel)
}
