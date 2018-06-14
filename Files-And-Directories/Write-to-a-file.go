package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	str := "this is a string"
	bytes := []byte(str)

	// using ioutil.WriteFile()
	err1 := ioutil.WriteFile("/tmp/file1", bytes, 0644)
	if err1 == nil {
		fmt.Println("File 1: Success")
	}

	// using file.Write()
	f2, err2a := os.Create("/tmp/file2")
	if err2a == nil {
		n2, err2b := f2.Write(bytes)
		if err2b == nil {
			fmt.Printf("File 2: Wrote %d bytes\n", n2)
		}
	}
	f2.Close()

	// using file.WriteString()
	f3, err3a := os.Create("/tmp/file3")
	if err3a == nil {
		n3, err3b := f3.WriteString(str)
		if err3b == nil {
			fmt.Printf("File 3: Wrote %d bytes\n", n3)
		}
	}
	f3.Close()
}
