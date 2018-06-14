package main

import (
   "fmt"
   "math"
)

func surface_area(radius,height float64) float64 {
   return math.Pi * radius * (radius + math.Sqrt(math.Pow(height,2) + math.Pow(radius,2)))
}

func main() {
   fmt.Println(surface_area(2.0, 5.0))
}


