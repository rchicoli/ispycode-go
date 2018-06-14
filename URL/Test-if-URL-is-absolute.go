package main

import (
	"fmt"
	"net/url"
)

func main() {

	str1 := "http://unknowhost.com/absolute/path"
	u1, _ := url.Parse(str1)
	fmt.Println("absolute:", u1.IsAbs())

	str2 := "/relitive/path"
	u2, _ := url.Parse(str2)
	fmt.Println("absolute:", u2.IsAbs())
}
