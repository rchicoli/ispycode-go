package main

import "log"

func function1() {
	function2()
}

func function2() {
	function3()
}

func function3() {
	log.Panic("PANIC")
}

func main() {
	function1()
}
