package main

import (
   "fmt"
   "math"
)

func volume (radius,height float64) float64 {
   return math.Pi * math.Pow(radius,2) * height /3
}

func main() {
   fmt.Println(volume(3.0,4.0))
}


