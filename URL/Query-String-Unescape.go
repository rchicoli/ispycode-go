package main

import (
	"fmt"
	"net/url"
)

func main() {

	// <abc?&%<
	str := "%3Cabc%3F%26%25%3E"

	newstr, _ := url.QueryUnescape(str)
	fmt.Println(newstr)
}
