package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://unknowhost.com/person.php?name=george&id=100"

	u, _ := url.Parse(str)

	values, _ := url.ParseQuery(u.RawQuery)

	name := values.Get("name")
	fmt.Println("name:", name)

	id := values.Get("id")
	fmt.Println("id:", id)

	unknown := values.Get("unknown")
	fmt.Println("unknown:", unknown)
}
