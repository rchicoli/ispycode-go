package main

import (
	"fmt"
	"net/url"
)

func main() {

	str1 := "http://ispycode.com/GO/URL/Get-Request-URI"
	u1, _ := url.Parse(str1)

	str2 := "/this/is/a/new/path"
	u2, _ := url.Parse(str2)

	u3 := u1.ResolveReference(u2)
	fmt.Println("url:", u3)
}
