package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"aaa", "bbb"}
	slice2 := []string{"ccc", "ddd"}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}
