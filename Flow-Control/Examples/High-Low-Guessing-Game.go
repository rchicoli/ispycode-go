package main

import (
	"fmt"
	"math/rand"
	"time"
)

func xrand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	mynumber := xrand(1, 100)
	attempts := 0
	var guess int

	fmt.Println("Guess a number between 1 and 100.")

	for attempts < 10 {
		fmt.Println("Take a guess:")
		fmt.Scanf("%d", &guess)
		attempts++
		if guess < mynumber {
			fmt.Println("Too low.")
		}
		if guess > mynumber {
			fmt.Println("Too high.")
		}
		if guess == mynumber {
			break
		}
	}
	if guess == mynumber {
		fmt.Printf("You guessed the number in %d tries\n", attempts)
	} else {
		fmt.Printf("The number was %d\n", mynumber)
	}
}
