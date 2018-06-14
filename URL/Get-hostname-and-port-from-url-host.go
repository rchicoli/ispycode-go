package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	str := "http://www.ispycode.com:80"

	u, _ := url.Parse(str)

	host, port, _ := net.SplitHostPort(u.Host)

	fmt.Println("host:", host)
	fmt.Println("port:", port)
}
