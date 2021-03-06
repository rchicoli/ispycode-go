package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {

	for i := 1; i < 65535; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial("tcp", "127.0.0.1:"+port)
		if err == nil {
			fmt.Println("Port", i, "open")
			conn.Close()
		}
	}
}
