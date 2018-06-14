package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(111111)
	b := big.NewInt(222222)

	fmt.Println(a)
	fmt.Println(b)

	a.Add(a, b)
	fmt.Println(a)
}
