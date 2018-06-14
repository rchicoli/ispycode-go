package main

import "fmt"
import "io/ioutil"

func main() {
	tmpdir, _ := ioutil.TempDir("", "mytempdir")
	fmt.Println(tmpdir)
}
