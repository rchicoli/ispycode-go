package main

import (
	"fmt"
	"math"
)

func main() {

	radius := 10
	diameter := (radius * 2) + 1
	fudge := .8

	for y := 1; y <= diameter; y++ {
		for x := 1; x <= diameter; x++ {
			a := math.Abs(float64(radius - x + 1))
			b := math.Abs(float64(radius - y + 1))
			c := int(math.Floor(math.Sqrt((a*a)+(b*b)) + fudge))
			if c > radius {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
