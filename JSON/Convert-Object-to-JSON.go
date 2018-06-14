package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Id   int
	Name string
}

func main() {

	employee := Employee{
		Id:   100,
		Name: "Tom Jones",
	}

	b, err := json.Marshal(employee)

	if err == nil {
		s := string(b)
		fmt.Println(s)
	}
}
