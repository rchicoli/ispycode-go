package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
)

func main() {
	misc, _ := load.Misc()
	fmt.Println("Running:", misc.ProcsRunning)
	fmt.Println("Blocked:", misc.ProcsBlocked)
}
