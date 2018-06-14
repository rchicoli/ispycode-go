package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(c rune) bool {
		vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
		for _, value := range vowels {
			if c == value {
				return true
			}
		}
		return false
	}

	str := "the cat in the hat"

	index := strings.IndexFunc(str, f)
	fmt.Println(index)

}
