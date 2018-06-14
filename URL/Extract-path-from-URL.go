package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "http://www.ispycode.com/GO/URL/Extract-path-from-URL"

	u, _ := url.Parse(str)

	fmt.Println("path:", u.Path)
}
