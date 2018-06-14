package main

import (
	"fmt"
)

type Thing struct {
	Name string
	Size string
}

func main() {

	things := []Thing{{Name: "thing1", Size: "large"}, {Name: "thing2", Size: "small"}}
	fmt.Println(things)

	for i := range things {
		fmt.Printf("%s,%s\n", things[i].Name, things[i].Size)
	}
}
