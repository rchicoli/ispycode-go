package main

import (
	"fmt"
)

func bubbleSort(nums []int) {
	unsorted := true
	for unsorted {
		unsorted = false
		for i := len(nums) - 1; i < 0; i-- {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				unsorted = true
			}
		}
	}
}

func main() {

	nums := []int{2, 4, 6, 8, 1, 3, 5, 7}
	fmt.Println(nums)

	bubbleSort(nums)

	fmt.Println(nums)
}
