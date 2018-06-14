package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	var tv syscall.Timeval
	if err := syscall.Gettimeofday(&tv); err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("Seconds: ", tv.Sec)
	fmt.Println("Micro:   ", tv.Usec)
	fmt.Println(time.Unix(tv.Sec, tv.Usec))
}
