package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{1, 3, 5, 7, 2, 4, 6, 8}

	b := sort.IntsAreSorted(i)

	fmt.Println(i)
	fmt.Println("Sorted:", b)

	sort.Ints(i)
	b = sort.IntsAreSorted(i)

	fmt.Println(i)
	fmt.Println("Sorted:", b)
}
