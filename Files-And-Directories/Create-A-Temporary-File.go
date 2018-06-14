package main

import "fmt"
import "io/ioutil"

func main() {
	tmpfile, _ := ioutil.TempFile("", "mytempfile")
	fmt.Println(tmpfile.Name())
}
