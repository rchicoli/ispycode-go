package main

import "fmt"

func main() {

	s := []int{111, 222, 333}
	fmt.Println(s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println(c)
}
