package main

import "fmt"

func main() {
	list := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("before:", list)
	gnomeSort(list)
	fmt.Println("after: ", list)
}

func gnomeSort(list []int) {
	for i, j := 1, 2; i < len(list); {
		if list[i-1] > list[i] {
			list[i-1], list[i] = list[i], list[i-1]
			i--
			if i > 0 {
				continue
			}
		}
		i = j
		j++
	}
}
