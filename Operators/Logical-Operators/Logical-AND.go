package main

import "fmt"

func main() {

	a := true
	b := false
	fmt.Println(a && a)
	fmt.Println(a && b)
	fmt.Println(b && a)
	fmt.Println(b && b)
}
