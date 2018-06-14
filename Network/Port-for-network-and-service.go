package main

import (
	"fmt"
	"net"
)

func main() {
	port, err := net.LookupPort("tcp", "ftp")
	if err == nil {
		fmt.Println("ftp port: ", port)
	}
}
