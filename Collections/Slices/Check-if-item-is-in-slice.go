package main

import "fmt"

func main() {

	s := []string{"apple", "grape", "plum"}
	fmt.Println(s)

	b := contains(s, "apple")
	fmt.Println(b)

	b = contains(s, "orange")
	fmt.Println(b)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
