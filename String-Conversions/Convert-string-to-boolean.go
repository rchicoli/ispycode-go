package main

import (
	"fmt"
	"strconv"
)

func main() {

	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Printf("Type: %T \n", b)
		fmt.Println("Value:", b)
	}

}
