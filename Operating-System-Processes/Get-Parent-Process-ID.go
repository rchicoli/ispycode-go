package main

import (
	"fmt"
	"os"
)

func main() {
	ppid := os.Getppid()
	fmt.Println("PPID:", ppid)
}
