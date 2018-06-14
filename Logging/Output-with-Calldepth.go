package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)

	log.Output(1, "One")
	log.Output(2, "Two")
	log.Output(3, "Three")
}
