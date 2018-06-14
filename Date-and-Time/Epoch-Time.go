package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now().Unix()
	fmt.Printf("Epoch time: %v\n", time)
}
