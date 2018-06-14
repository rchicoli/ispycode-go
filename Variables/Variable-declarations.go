package main

import "fmt"

func main() {

	// int
	var number = 50
	fmt.Println(number)

	// string
	var str = "Hello"
	fmt.Println(str)

	// boolean
	var b = true
	fmt.Println(b)

	// struct
	type person struct {
		first string
		last  string
	}

	var p = new(person)
	p.first = "tom"
	p.last = "brady"
	fmt.Println(*p)
}
