opackage main

import  "fmt"

func surface_area(length,width,height float64) float64 {
   return 2.0 * (width*length + height*length + height*width)
}

func main() {
   fmt.Println(surface_area(2.0, 3.0, 4.0))
}


