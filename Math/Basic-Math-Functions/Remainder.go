package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10.0
	y := 4.1

	z := math.Remainder(x, y)
	fmt.Println(z)
}
