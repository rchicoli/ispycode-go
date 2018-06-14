package main

import "fmt"

func vals() (int, string) {
	return 3, "hello"
}
func main() {
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
