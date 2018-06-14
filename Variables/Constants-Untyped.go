package main

import (
	"fmt"
	"reflect"
)

func main() {

	const x = 50
	fmt.Println(reflect.TypeOf(x), x)

	const y = int64(50)
	fmt.Println(reflect.TypeOf(y), y)

}
