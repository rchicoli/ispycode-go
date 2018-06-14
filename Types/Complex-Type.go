package main

import "fmt"

func main() {

	a := complex(100, 8)
	fmt.Println(a)

	fmt.Println(real(a))

	fmt.Println(imag(a))

	fmt.Printf("Type:%T \n", a)
}
