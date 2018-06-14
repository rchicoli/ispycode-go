package main

import (
   "fmt"
   "math"
)

func volume (radius float64) float64 {
   return 4.0/3.0 * math.Pi * math.Pow(radius,3)
}

func main() {
   fmt.Println(volume(3.0))
}


