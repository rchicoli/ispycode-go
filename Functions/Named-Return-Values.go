package main

import "fmt"

func swap(a int, b int) (x, y int) {
	x = b
	y = a
	return
}

func main() {
	x := 1
	y := 2
	x, y = swap(x, y)
	fmt.Println(x)
	fmt.Println(y)
}
