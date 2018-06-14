package main

import (
   "fmt"
   "math"
)

func surface_area(length,width,height float64) float64 {
    return length * width +
	   length * math.Sqrt(math.Pow((width/2.0),2) + math.Pow(height,2)) +
	   width * math.Sqrt(math.Pow((length/2.0),2) + math.Pow(height,2))
}

func main() {
   fmt.Println(surface_area(3.0,4.0,5.0))
}


