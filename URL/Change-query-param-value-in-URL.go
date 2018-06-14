package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://unknowhost.com/person.php?name=george&id=100"

	u, _ := url.Parse(str)
	fmt.Println("url:", u)

	values, _ := url.ParseQuery(u.RawQuery)
	values.Set("name", "tom")

	u.RawQuery = values.Encode()
	fmt.Println("new url:", u)

}
