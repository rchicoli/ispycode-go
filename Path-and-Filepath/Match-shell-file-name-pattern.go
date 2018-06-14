package main

import (
	"fmt"
	"path"
)

func main() {

	mypath := "/Users/dennis/ISPY/GO/example.go"

	results, _ := path.Match("/*/*/*/*/*.go", mypath)
	fmt.Println("Match:", results)

	results, _ = path.Match("/*/*/*/*/exampl?.go", mypath)
	fmt.Println("Match:", results)
}
