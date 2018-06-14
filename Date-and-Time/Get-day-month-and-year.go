package main

import (
	"fmt"
	"time"
)

func main() {

	// Get the current time.
	t := time.Now()

	// Year returns the year
	y := t.Year()

	// Month returns the month of the year
	m := t.Month()

	// Day returns the day of the month
	d := t.Day()

	fmt.Println("Year :", y)
	fmt.Println("Month:", m)
	fmt.Println("Day  :", d)
}
