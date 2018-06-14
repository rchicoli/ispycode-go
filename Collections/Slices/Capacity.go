package main

import "fmt"

func main() {

	// create slice length 1, capacity 5
	s := make([]int, 1, 5)
	s[0] = 99

	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	// append 3 items
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))

	// append 3 more items
	s = append(s, 4, 5, 6)
	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Capacity:", cap(s))
}
