package main

import "fmt"

func main() {

	i := 1234
	f := 1.23
	b := true
	p := &i
	s := []int{1, 2, 3, 4}
	str := "Hello"

	fmt.Printf("The Go-syntax representation of the value %#v \n", i)
	fmt.Printf("The Go-syntax representation of the value %#v \n", f)
	fmt.Printf("The Go-syntax representation of the value %#v \n", b)
	fmt.Printf("The Go-syntax representation of the value %#v \n", p)
	fmt.Printf("The Go-syntax representation of the value %#v \n", s)
	fmt.Printf("The Go-syntax representation of the value %#v \n", str)
}
