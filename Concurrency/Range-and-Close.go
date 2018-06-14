package main

import "fmt"

func main() {
	queue := make(chan string, 3)

	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	// iterate over each element in the queue.
	for elem := range queue {
		fmt.Println(elem)
	}
}
