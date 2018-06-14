package main

import (
	"fmt"
	"time"
)

func main() {

	t0 := time.Now()

	// sleep 2 seconds
	time.Sleep(2 * time.Second)

	t1 := time.Now()

	d := t1.Sub(t0)

	fmt.Println("Duration:", d)

	// duration in seconds
	fmt.Println("Duration:", d.Seconds())

	// duration in hours
	fmt.Println("Duration:", d.Hours())
}
