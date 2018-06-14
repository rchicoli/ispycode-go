package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	tom := person{name: "Tom", age: 50}

	fmt.Println(tom)
	fmt.Println(tom.name)
	fmt.Println(tom.age)
}
