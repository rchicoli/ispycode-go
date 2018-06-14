package main

import (
	"fmt"
	"time"
)

func main() {

	// Get the current time.
	t := time.Now()

	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	fmt.Println("Hour  :", h)
	fmt.Println("Minute:", m)
	fmt.Println("Second:", s)
}
