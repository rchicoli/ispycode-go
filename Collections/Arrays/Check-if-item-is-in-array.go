package main

import "fmt"

func main() {

	arr := [3]string{}
	arr[0] = "apple"
	arr[1] = "plum"
	arr[2] = "orange"
	fmt.Println(arr)

	b := contains(arr, "plum")
	mt.Println(b)

	b = contains(arr, "grape")
	fmt.Println(b)
}

func contains(arr [3]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
