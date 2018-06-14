package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	// print current time
	fmt.Println(time.Now())

	// Command that sleeps for 10 seconds
	cmd := exec.Command("sleep", "10")

	// start command
	cmd.Start()

	// wait for command to finish
	cmd.Wait()

	// print current time
	fmt.Println(time.Now())

}
