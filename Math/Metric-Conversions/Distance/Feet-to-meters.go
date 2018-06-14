package main

import "fmt"

func feet2meters(feet float64) float64 {
	return feet * 0.3048
}

func main() {
	fmt.Println(feet2meters(6.0))
}
