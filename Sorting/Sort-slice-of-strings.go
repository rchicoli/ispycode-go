package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"aa", "cc", "ee", "bb", "dd", "ff"}

	// tests whether a slice of strings is sorted in increasing order.
	b := sort.StringsAreSorted(s)

	fmt.Println(s)
	fmt.Println("Sorted:", b)

	// Strings sorts a slice of strings in increasing order.
	sort.Strings(s)

	b = sort.StringsAreSorted(s)

	fmt.Println(s)
	fmt.Println("Sorted:", b)
}
