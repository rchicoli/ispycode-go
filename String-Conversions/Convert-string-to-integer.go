package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "1234"

	// using strconv.Atoi
	i1, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println(i1)
	}

	// using strconv.ParseInt
	i2, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Println(i2)
	}

}
