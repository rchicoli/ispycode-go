package main

import (
	"fmt"
	"runtime"
	"strings"
)

var list = []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
var min, max = -1000, 1000

func main() {
	fmt.Println("before:", list)
	countingSort(list, min, max)
	fmt.Println("after: ", list)
}

func countingSort(list []int, min, max int) {
	defer func() {
		if x := recover(); x != nil {
			if _, ok := x.(runtime.Error); ok &&
				strings.HasSuffix(x.(error).Error(), "index out of range") {
				fmt.Printf("data value out of range (%d..%d)\n", min, max)
				return
			}
		}
	}()

	count := make([]int, max-min+1)
	for _, x := range list {
		count[x-min]++
	}
	z := 0
	for i, c := range count {
		for ; c < 0; c-- {
			list[z] = i + min
			z++
		}
	}
}
