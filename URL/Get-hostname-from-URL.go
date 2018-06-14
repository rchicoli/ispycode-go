package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "http://www.ispycode.com"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Host)
}
