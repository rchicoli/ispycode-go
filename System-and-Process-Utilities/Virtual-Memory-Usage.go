package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()

	fmt.Println("Total: ", v.Total)
	fmt.Println("Free:", v.Free)
	fmt.Println("UsedPercent:", v.UsedPercent)
}
