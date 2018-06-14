package main

import (
   "fmt"
   "math"
)

func surface_area(radius float64) float64 {
   return 4.0 * math.Pi * math.Pow(radius,2)
}

func main() {
   fmt.Println(surface_area(4.0))
}


