package main

import "fmt"

func main() {

	nums := []int{11, 22, 33, 44, 55}

	for i := range nums {
		fmt.Println(nums[i])
	}
}
