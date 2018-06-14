package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "http://unknowhost.com/person.php?first=george&last=washington"

	u, _ := url.Parse(str)
	fmt.Println("original:", u)

	q, _ := url.ParseQuery(u.RawQuery)

	fmt.Println("BEFORE")
	for key, value := range q {
		fmt.Println(key, ":", value)
	}

	fmt.Println()

	q.Add("age", "100")

	fmt.Println("AFTER")
	for key, value := range q {
		fmt.Println(key, ":", value)
	}

	u.RawQuery = q.Encode()

	fmt.Println("modified:", u)

}
