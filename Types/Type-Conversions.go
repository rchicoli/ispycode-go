package main

import (
	"fmt"
)

func main() {

	var i int = 42
	fmt.Printf("i:%v Type:%T \n", i, i)

	var f float64 = float64(i)
	fmt.Printf("f:%v Type:%T \n", f, f)

	var u uint = uint(i)
	fmt.Printf("u:%v Type:%T \n", u, u)
}
