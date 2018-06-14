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

	// String contains two JSON rows.
	text := `[{"Id": 100, "Name": "Tom"}, {"Id": 200, "Name": "Bob"}]`

	// Get byte slice from string.
	bytes := []byte(text)

	// Unmarshal string into structs.
	var employee []Employee
	json.Unmarshal(bytes, &employee)

	// Loop over structs and display them.
	for l := range employee {
		fmt.Printf("Id = %v, Name = %v", employee[l].Id, employee[l].Name)
		fmt.Println()
	}
}
