package main

import "fmt"
import "math"

func area(side float64) float64 {
   return (math.Sqrt(3.0)/4.0 * math.Pow(side,2))
}

func main() {
   fmt.Println(area(5.0))
}


