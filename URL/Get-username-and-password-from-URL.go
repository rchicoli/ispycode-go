package main

import (
	"fmt"
	"net/url"
)

func main() {

	str := "https://tom.jones:mypass@unknowhost.com/location/mystuff.git"

	u, _ := url.Parse(str)

	// username
	fmt.Println("user:", u.User.Username())

	//password
	p, _ := u.User.Password()
	fmt.Println("pass:", p)
}
