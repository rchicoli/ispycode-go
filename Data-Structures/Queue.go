package main

import (
	"errors"
	"fmt"
)

type queue struct {
	h, count, size int
	buf            []interface{}
}

func NewQueue(maxsize int) *queue {
	return &queue{
		h:     0,
		count: 0,
		buf:   make([]interface{}, maxsize),
		size:  maxsize,
	}
}

func (this *queue) Push(data interface{}) error {
	if this.count == this.size {
		return errors.New("queue full")
	}
	this.buf[(this.h+this.count)%this.size] = data
	this.count++
	return nil
}

func (this *queue) Pop() (interface{}, error) {
	if rect, err := this.Peek(); err == nil {
		this.h++
		if this.h == this.size {
			this.h = 0
		}
		this.count--
		return rect, err
	} else {
		return rect, err
	}
}

func (this *queue) Peek() (interface{}, error) {
	if this.count == 0 {
		return nil, errors.New("queue empty")
	}
	return this.buf[this.h], nil
}

func (this *queue) Length() int {
	return this.count
}

func main() {
	q := NewQueue(10)
	q.Push("aaa")
	q.Push("bbb")
	q.Push("ccc")

	for {
		str, err := q.Pop()
		if err == nil {
			fmt.Println(str)
		} else {
			fmt.Println("Error:", err)
			break
		}
	}
}
