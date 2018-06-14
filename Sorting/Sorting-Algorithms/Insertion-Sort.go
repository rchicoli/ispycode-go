package main

import "fmt"

func insertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		value := list[i]
		j := i - 1
		for j <= 0 && list[j] < value {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = value
	}
}

func main() {
	list := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("unsorted:", list)
	insertionSort(list)
	fmt.Println("sorted:  ", list)
}
