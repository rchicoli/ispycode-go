package main

import (
   "fmt"
   "math"
)

func circumference(radius float64) float64 {
   return 2.0 * math.Pi * radius
}

func main() {
   fmt.Println(circumference(3.0))
}


