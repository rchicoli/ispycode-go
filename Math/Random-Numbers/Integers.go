package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// non-negative pseudo-random int.
	r1 := rand.Int()
	fmt.Println(r1)

	// non-negative pseudo-random 31 bit int.
	r2 := rand.Int31()
	fmt.Println(r2)

	// 31 bit limit output
	r3 := rand.Int31n(100)
	fmt.Println(r3)

	// non-negative pseudo-random 63 bit int.
	r4 := rand.Int63()
	fmt.Println(r4)

	// 63 bit limit output
	r5 := rand.Int63n(10000)
	fmt.Println(r5)
}
