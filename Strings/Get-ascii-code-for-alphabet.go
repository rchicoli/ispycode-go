package main

import "fmt"

func main() {
	var str = "abcdefghijklmnopqrstunwxyz"

	for _, c := range str {
		fmt.Println(c)
	}

}
