package main

import (
	"fmt"
	"net"
)

func main() {
	host, err := net.LookupAddr("50.62.227.1")
	if err == nil {
		fmt.Println(host)
	}
}
