package main

import "fmt"

func updateValue(val *int) {
	*val = *val + 100
}

func main() {
	x := 1000
	updateValue(&x)
	fmt.Println("x:", x)
}
