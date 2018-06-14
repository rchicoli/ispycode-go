package main

import "fmt"

type Pet interface {
	speak()
}

func speak(p Pet) {
	p.speak()
}

type Cat struct {
	Pet
}

func (c Cat) speak() {
	fmt.Println("I am a cat.")
}

type Dog struct {
	Pet
}

func (d Dog) speak() {
	fmt.Println("I am a Dog.")
}

func main() {

	c := &Cat{}
	speak(c)

	d := &Dog{}
	speak(d)

}
