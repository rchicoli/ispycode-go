package main

import (
	"fmt"
	"os/exec"
)

func main() {

	// Command returns the Cmd struct to execute the named program.
	cmd := exec.Command("/bin/date")

	// Output runs the command and returns its standard output.
	output, err := cmd.Output()
	if err == nil {
		fmt.Println(string(output))
	}
}
