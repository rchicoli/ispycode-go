package main

import "fmt"

func area(base, height float64) float64 {
	return (base * height) / 2.0
}

func main() {
	fmt.Println(area(5.0, 4.0))
}
