package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("myfile.txt")

	if err != nil {
		fmt.Print(err)
	}

	// convert content to a 'string'
	str := string(b)

	// print the content as a 'string'
	fmt.Println(str)
}
