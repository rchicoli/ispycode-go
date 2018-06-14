package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	cpu, _ := cpu.Info()

	for i, v := range cpu {
		fmt.Println("CPU:", i)
		fmt.Println(" Model    : ", v.ModelName)
		fmt.Println(" Cores    : ", v.Cores)
		fmt.Println(" CacheSize: ", v.CacheSize)
	}
}
