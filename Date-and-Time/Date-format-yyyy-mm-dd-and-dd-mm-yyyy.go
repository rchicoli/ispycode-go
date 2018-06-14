package main

import (
   "fmt"
   "time"
)

func main() {

   now := time.Now()
   fmt.Println("Now :", now)

   fmt.Println("mm-dd-yyyy date format : ", now.Format("01-02-2006"))
   fmt.Println("yyyy-mm-dd date format : ", now.Format("2006-01-02"))


