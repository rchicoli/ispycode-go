package main

import "fmt"

func removeDuplicates(a []int) []int {

	mymap := map[int]bool{}
	result := []int{}

	for v := range a {
		if mymap[a[v]] == true {
		} else {
			mymap[a[v]] = true
			result = append(result, a[v])
		}
	}
	return result
}

func main() {
	i := []int{1, 2, 3, 4, 1, 2, 3, 1, 2, 1}
	fmt.Println(i)
	i = removeDuplicates(i)
	fmt.Println(i)
}
