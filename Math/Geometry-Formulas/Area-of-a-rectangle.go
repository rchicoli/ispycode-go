package main

import "fmt"

func area(width,length float64) float64 {
   return length * width
}

func main() {
   fmt.Println(area(3.0,4.0))
}


