package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://unknowhost.com/person.php?first=george&last=washington&age=100"

	u, _ := url.Parse(str)

	m, _ := url.ParseQuery(u.RawQuery)

	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
