package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	priority int
	name     string
}

type TaskPQ []Task

func (self TaskPQ) Len() int { return len(self) }
func (self TaskPQ) Less(i, j int) bool {
	return self[i].priority < self[j].priority
}
func (self TaskPQ) Swap(i, j int)       { self[i], self[j] = self[j], self[i] }
func (self *TaskPQ) Push(x interface{}) { *self = append(*self, x.(Task)) }
func (self *TaskPQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}

func main() {
	pq := &TaskPQ{
		{3, "Vacuum"},
		{2, "Feed cat"},
		{4, "Arrange play date"},
		{1, "Pack for trip"}}

	// heapify
	heap.Init(pq)

	// enqueue
	heap.Push(pq, Task{2, "Iron"})

	for pq.Len() != 0 {
		// dequeue
		fmt.Println(heap.Pop(pq))
	}
}
