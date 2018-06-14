package main

import (
	"fmt"
	"math"
)

func surface_area(edge float64) float64 {
	return 6.0 * math.Pow(edge, 2)
}

func main() {
	fmt.Println(surface_area(3.0))
}
