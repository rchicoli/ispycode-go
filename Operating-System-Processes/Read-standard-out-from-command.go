package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// Command returns the Cmd struct to execute the named program.
	cmd := exec.Command("date")

	// pipe that will be connected to the command's standard output
	stdout, _ := cmd.StdoutPipe()

	// start command
	cmd.Start()

	// read standard output
	buf, _ := ioutil.ReadAll(stdout)

	// print standard output
	fmt.Println("stdout:", string(buf))
}
