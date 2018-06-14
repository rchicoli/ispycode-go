package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://unknowhost.com"

	u, _ := url.Parse(str)
	fmt.Println("orig url    :", u)

	v := url.Values{}
	v.Add("name", "george")
	v.Add("age", "100")

	encoded := v.Encode()
	fmt.Println("query string:", encoded)

	u.RawQuery = encoded
	fmt.Println("new url     :", u)
}
