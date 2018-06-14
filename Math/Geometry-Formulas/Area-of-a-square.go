package main

import (
	"fmt"
	"math"
)

func area(length float64) float64 {
	return math.Pow(length, 2)
}

func main() {
	fmt.Println(area(10.0))
}
