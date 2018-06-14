package main

import (
   "fmt"
   "log"
)

func main() {
   fmt.Println("Before fatal error")

   log.Fatalln("Died here")

   fmt.Println("After fatal error
}


