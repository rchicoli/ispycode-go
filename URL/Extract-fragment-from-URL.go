package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "https://en.wikipedia.org/wiki/Fragment_identifier#Examples"

	u, _ := url.Parse(str)

	fmt.Println("fragment:", u.Fragment)
}
