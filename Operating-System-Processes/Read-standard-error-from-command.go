package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// Command returns the Cmd struct to execute the named program.
	cmd := exec.Command("/bin/ls", "no_dir")

	// pipe that will be connected to the command's standard error
	stderr, _ := cmd.StderrPipe()

	// start command
	cmd.Start()

	// read standard error
	buf, _ := ioutil.ReadAll(stderr)

	// print standard error
	fmt.Println("stderr:", string(buf))
}
