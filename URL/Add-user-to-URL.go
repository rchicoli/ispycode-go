package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "https://unknowhost.com/location/mystuff.git"

	u, _ := url.Parse(str)
	fmt.Println("original:", u)

	user := "tom.jones"
	pass := "secret"
	u.User = url.UserPassword(user, pass)

	fmt.Println("new url:", u)
}
