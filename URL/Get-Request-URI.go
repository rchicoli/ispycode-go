package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://ispycode.com/GO/URL/Get-Request-URI"
	u, _ := url.Parse(str)
	uri := u.RequestURI()
	fmt.Println("uri:", uri)
}
