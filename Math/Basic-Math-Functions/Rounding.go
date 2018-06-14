package main

import (
	"fmt"
	"math"
)

func main() {

	f := 1.0
	for i := 0; i < 10; i++ {
		f = f + .1
		round := math.Floor(f + .5)
		fmt.Printf("%.1f : %f\n", f, round)
	}
}
