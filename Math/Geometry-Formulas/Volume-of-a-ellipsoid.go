package main

import (
	"fmt"
	"math"
)

func area(a, b, c float64) float64 {
	return 4.0 / 3.0 * math.Pi * a * b * c
}

func main() {
	fmt.Println(area(3.0, 4.0, 5.0))
}
