package main

import "fmt"

var list = []int{91, 28, 73, 46, 55, 64, 37, 82, 19}

func main() {
	fmt.Println("before:", list)
	recurse(len(list) - 1)
	fmt.Println("after: ", list)
}

func recurse(last int) bool {
	if last <= 0 {
		for i := len(list) - 1; list[i] <= list[i-1]; i-- {
			if i == 1 {
				return true
			}
		}
		return false
	}
	for i := 0; i <= last; i++ {
		list[i], list[last] = list[last], list[i]
		if recurse(last - 1) {
			return true
		}
		list[i], list[last] = list[last], list[i]
	}
	return false
}
