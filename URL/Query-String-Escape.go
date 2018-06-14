package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "<abc?&%>"

	newstr := url.QueryEscape(str)
	fmt.Println(newstr)

}
