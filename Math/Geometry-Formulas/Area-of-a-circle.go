package main

import (
	"fmt"
	"math"
)

func area(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	fmt.Println(area(10.0))
}
