package main

import "fmt"

func getFunction() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := getFunction()

	x := nextInt()
	fmt.Println(x)

	x = nextInt()
	fmt.Println(x)
}
