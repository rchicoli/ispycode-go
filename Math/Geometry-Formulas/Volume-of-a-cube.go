package main

import "fmt"
import "math"

func volume(edge float64) float64 {
	return math.Pow(edge, 3)
}

func main() {
	fmt.Println(volume(5.0))
}
