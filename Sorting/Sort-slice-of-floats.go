package main

import (
	"fmt"
	"sort"
)

func main() {

	f := []float64{1.1, 3.3, 5.5, 2.2, 4.4, 6.6}

	// tests whether float64s are sorted in increasing order.
	b := sort.Float64sAreSorted(f)

	fmt.Println(f)
	fmt.Println("Sorted:", b)

	// sorts float64s in increasing order.
	sort.Float64s(f)

	b = sort.Float64sAreSorted(f)

	fmt.Println(f)
	fmt.Println("Sorted:", b)
}
