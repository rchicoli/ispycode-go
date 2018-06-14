package main

import "fmt"

func volume (length,width,height float64) float64 {
   return length * width * height / 3.0
}

func main() {
   fmt.Println(volume(3.0,4.0,5.0))
}


