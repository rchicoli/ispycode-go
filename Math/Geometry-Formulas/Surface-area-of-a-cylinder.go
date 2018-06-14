package main

import (
   "fmt"
   "math"
)

func surface_area(radius,height float64) float64 {
   return 2.0 * math.Pi * radius * height + 2.0 * math.Pi * math.Pow(radius,2)
}

func main() {
   fmt.Println(surface_area(2.0, 5.0))
}


