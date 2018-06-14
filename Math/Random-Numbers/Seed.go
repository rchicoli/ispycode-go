package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	now := time.Now()

	rand.Seed(now.Unix())

	r1 := rand.Int31()
	fmt.Println(r1)

}
