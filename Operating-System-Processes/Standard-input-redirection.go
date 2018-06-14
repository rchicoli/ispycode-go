package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// Command returns the Cmd struct to execute the named program.
	cmd := exec.Command("wc", "-m")

	// pipe that will be connected to the command's standard input
	stdin, _ := cmd.StdinPipe()

	// pipe that will be connected to the command's standard output
	stdout, _ := cmd.StdoutPipe()

	// start command
	cmd.Start()

	// write hello world to standard input
	stdin.Write([]byte("hello world"))
	stdin.Close()

	// read standard output
	buf, _ := ioutil.ReadAll(stdout)

	// print standard error output
	fmt.Println("stdout:", string(buf))
}
