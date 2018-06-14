package main

import (
	"fmt"
	"time"
)

func main() {
	list := []uint{90, 10, 80, 20, 70, 30, 60, 40, 50}
	fmt.Println("before:", list)
	sorted := sleepsort(list)
	fmt.Println("after: ", sorted)
}

func sleepsort(list []uint) []uint {
	sorted := []uint{}
	out := make(chan uint)
	for _, i := range list {
		go func(n uint) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			out <- n
		}(i)
	}
	for _ = range list {
		sorted = append(sorted, <-out)
	}
	return (sorted)
}
