package main

import "fmt"

func main() {
	list := []int{90, -80, 70, -60, 50, -40, 30, -20, 10, 0}
	fmt.Println("before:", list)
	stoogesort(list)
	fmt.Println("after: ", list)
}

func stoogesort(list []int) {
	last := len(list) - 1
	if list[last] < list[0] {
		list[0], list[last] = list[last], list[0]
	}
	if last > 1 {
		t := len(list) / 3
		stoogesort(list[:len(list)-t])
		stoogesort(list[t:])
		stoogesort(list[:len(list)-t])
	}
}
