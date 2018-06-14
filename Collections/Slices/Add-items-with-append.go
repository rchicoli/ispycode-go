package main

import "fmt"

func main() {

	a := []int{111, 222, 333}
	fmt.Println(a)

	a = append(a, 444)
	fmt.Println(a)

	a = append(a, 555, 666, 777)
	fmt.Println(a)
}
