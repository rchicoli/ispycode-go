package main

import (
	"fmt"
	"syscall"
)

func main() {

	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err != nil {
		fmt.Println("Error:", err)
	}
	// uptime seconds since boot
	fmt.Println("uptime:", info.Uptime)

	// 1, 5, and 15 minute load averages
	fmt.Println("loads:", info.Loads)

	// Total usable main memory size
	fmt.Println("total ram:", info.Totalram)

	// Available memory size
	fmt.Println("free ram:", info.Freeram)

	// Amount of shared memory
	fmt.Println("shared ram:", info.Sharedram)

	// Total swap space size
	fmt.Println("total swap:", info.Totalswap)

	// Swap space still available
	fmt.Println("free swap:", info.Freeswap)
}
