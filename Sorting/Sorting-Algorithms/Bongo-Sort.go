package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	list := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("unsorted:", list)

	temp := make([]int, len(list))
	copy(temp, list)
	for !sort.IntsAreSorted(temp) {
		for i, v := range rand.Perm(len(list)) {
			temp[i] = list[v]
		}
	}

	fmt.Println("sorted  :", temp)
}
