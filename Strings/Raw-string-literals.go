package main

import (
	"fmt"
)

func main() {

	str1 := "Big cat \nLittle dog"
	fmt.Println(str1)

	// The newline sequence is treated as raw chars.
	str2 := `Big cat \n Little dog`
	fmt.Println(str2)

}
