package main

import (
	"fmt"
	"math/rand"
)

func main() {

	s := rand.Perm(10)
	fmt.Println(s)
}
