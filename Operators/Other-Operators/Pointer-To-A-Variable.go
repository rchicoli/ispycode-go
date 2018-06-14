package main

import "fmt"

func main() {

	var i = 100
	var address = &i
	var value = *address
	fmt.Println(value)
}
