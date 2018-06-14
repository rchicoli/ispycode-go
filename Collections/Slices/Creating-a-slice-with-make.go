package main

import "fmt"

func main() {

	a := make([]int, 5) // len(a)=5
	fmt.Println("Length:", len(a))
	fmt.Println("Capacity:", cap(a))

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println("Length:", len(b))
	fmt.Println("Capacity:", cap(b))
}
