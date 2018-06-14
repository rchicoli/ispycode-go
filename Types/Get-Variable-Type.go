package main

import (
	"fmt"
	"reflect"
)

func main() {

	// integer
	i := 1

	// float
	f := 1.23456

	// string
	s := "Hello World"

	fmt.Println("i type: ", reflect.TypeOf(i))
	fmt.Println("f type: ", reflect.TypeOf(f))
	fmt.Println("s type: ", reflect.TypeOf(s))
}
