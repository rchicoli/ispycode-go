package main

import "fmt"

func myadd(a int, b int) int {
	return a + b
}

func main() {
	x := 10
	y := 5
	z := myadd(x, y)
	fmt.Println(z)
}
