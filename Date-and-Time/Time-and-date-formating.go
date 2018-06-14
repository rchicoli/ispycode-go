package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	// month day year
	fmt.Println(t.Format("January 1 2006"))

	// hour minute second
	fmt.Println(t.Format("03:04:05"))

	// month day year hour minute PM timezine
	fmt.Println(t.Format("Jan 01 2016 03:04:05 PM MST"))
}
