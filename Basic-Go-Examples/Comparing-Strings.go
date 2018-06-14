package main

import "fmt"

func main() {

	str1 := "hello"
	str2 := "hello"
	str3 := "goodbye"

	if str1 == str2 {
		fmt.Println("str1 == str2")
	}

	if str1 != str3 {
		fmt.Println("str1 != str3")
	}
}
