package main

import "fmt"

func main() {

	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
}
