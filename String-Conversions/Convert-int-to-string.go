package main

import (
	"fmt"
	"strconv"
)

func main() {

	i1 := 1234
	str1 := strconv.Itoa(i1)
	fmt.Println(str1)

	i2 := 5678
	str2 := strconv.FormatInt(int64(i2), 10)
	fmt.Println(str2)
}
