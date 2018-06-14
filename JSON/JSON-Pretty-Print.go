package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Employee struct {
	Id   int
	Name string
}

func main() {

	employees := []Employee{}
	for i := 1; i <= 5; i++ {
		employees = append(employees, Employee{i, "name" + strconv.Itoa(i)})
	}

	b, err := json.MarshalIndent(employees, "", " ")
	if err == nil {
		s := string(b)
		fmt.Println(s)
	}
}
