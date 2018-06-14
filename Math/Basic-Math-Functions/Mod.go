package main

import (
	"fmt"
	"math"
)

func main() {

	x := 10.0
	y := 2.9
	z := math.Mod(x, y)
	fmt.Println(z)
}
