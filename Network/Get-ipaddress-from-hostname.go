package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupIP("ispycode.com")
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}
}
