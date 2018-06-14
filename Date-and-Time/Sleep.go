package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())

	// sleep 500 Nanosecond
	time.Sleep(500 * time.Nanosecond)
	fmt.Println(time.Now())

	// sleep 500 Microseconds
	time.Sleep(500 * time.Microsecond)
	fmt.Println(time.Now())

	// sleep 500 Milliseconds
	time.Sleep(500 * time.Millisecond)
	fmt.Println(time.Now())

	// sleep 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())

	// sleep 5 minutes
	time.Sleep(5 * time.Minute)
	fmt.Println(time.Now())

	// sleep 2 hours
	time.Sleep(2 * time.Hour)
	fmt.Println(time.Now())
}
