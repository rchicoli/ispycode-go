package main

import "fmt"

func main() {
	list := []int{90, -80, 70, -60, 50, -40, 30, -20, 10, 0}
	fmt.Println("before:", list)
	selectionsort(list)
	fmt.Println("after: ", list)
}

func selectionsort(list []int) {
	last := len(list) - 1
	for i := 0; i < last; i++ {
		aMin := list[i]
		iMin := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < aMin {
				aMin = list[j]
				iMin = j
			}
		}
		list[i], list[iMin] = aMin, list[i]
	}
}
