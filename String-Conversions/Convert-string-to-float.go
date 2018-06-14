package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, err := strconv.ParseFloat("1.234", 64)
	if err == nil {
		fmt.Printf("Type: %T \n", f)
		fmt.Println("Value:", f)
	}

}
