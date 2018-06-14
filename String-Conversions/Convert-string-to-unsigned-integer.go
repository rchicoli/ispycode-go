package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.ParseUint("1234", 10, 64)
	if err == nil {
		fmt.Printf("Type: %T \n", i)
		fmt.Println("Value:", i)
	}

}
