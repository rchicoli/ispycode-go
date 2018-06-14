package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		fmt.Println("ID:", id, " count:", count)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
}
