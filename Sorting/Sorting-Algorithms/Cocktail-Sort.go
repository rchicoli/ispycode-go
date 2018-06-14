package main

import "fmt"

func main() {
	list := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("before:", list)
	cocktailSort(list)
	fmt.Println("after: ", list)
}

func cocktailSort(list []int) {
	last := len(list) - 1
	for {
		swapped := false
		for i := 0; i < last; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
		swapped = false
		for i := last - 1; i >= 0; i-- {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
