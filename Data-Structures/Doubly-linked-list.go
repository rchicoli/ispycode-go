package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist := list.New()
	e4 := mylist.PushBack(4)
	e1 := mylist.PushFront(1)
	mylist.InsertBefore(3, e4)
	mylist.InsertAfter(2, e1)

	// Iterate through the list and print its contents.
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
