package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func main() {
	uptime, _ := host.Uptime()
	fmt.Println("Total:", uptime, "seconds")

	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	fmt.Printf("%d days, %d hours, %d minutes", days, hours, minutes)
}
