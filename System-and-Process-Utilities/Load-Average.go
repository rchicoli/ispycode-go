package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
)

func main() {
	load, _ := load.Avg()
	fmt.Println(" 1 min ave:", load.Load1)
	fmt.Println(" 5 min ave:", load.Load5)
	fmt.Println("15 min ave:", load.Load15)

}
