package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// pseudo-random float32 betwwen 0.0 -< 1.0
	r1 := rand.Float32()
	fmt.Println(r1)

	// pseudo-random float64 betwwen 0.0 -< 1.0
	r2 := rand.Float64()
	fmt.Println(r2)

	// normally distributed float64
	r3 := rand.NormFloat64()
	fmt.Println(r3)

	// exponentially distributed float64
	r4 := rand.ExpFloat64()
	fmt.Println(r4)
}
