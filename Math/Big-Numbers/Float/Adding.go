package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(1.11111111)
	b := big.NewFloat(2.2222)

	fmt.Println(a)
	fmt.Println(b)

	a.Add(a, b)
	fmt.Println(a)
}
