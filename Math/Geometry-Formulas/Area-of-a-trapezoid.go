package main

import "fmt"

func area(a, b, h float64) float64 {
	return (a + b) / 2.0 * h
}

func main() {
	fmt.Println(area(3.0, 4.0, 5.0))
}
