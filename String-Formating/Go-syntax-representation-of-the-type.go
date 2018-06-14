package main

import "fmt"

func main() {

	i := 1234
	f := 1.23
	b := true
	p := &i
	s := []int{1, 2, 3, 4}
	str := "Hello"

	fmt.Printf("Go-syntax representation of the type %T \n", i)
	fmt.Printf("Go-syntax representation of the type %T \n", f)
	fmt.Printf("Go-syntax representation of the type %T \n", b)
	fmt.Printf("Go-syntax representation of the type %T \n", p)
	fmt.Printf("Go-syntax representation of the type %T \n", s)
	fmt.Printf("Go-syntax representation of the type %T \n", str)
}
