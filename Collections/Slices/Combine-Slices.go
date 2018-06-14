package main

import "fmt"

func main() {

	s1 := []int{111, 222, 333}
	fmt.Println(s1)

	s2 := []int{444, 555, 666}
	fmt.Println(s2)

	s1 = append(s1, s2...)
	fmt.Println(s1)
}
