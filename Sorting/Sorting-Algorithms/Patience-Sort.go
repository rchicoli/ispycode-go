package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Pile []int

func (self Pile) Top() int { return self[len(self)-1] }
func (self *Pile) Pop() int {
	x := (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return x
}

type PileHeap []Pile

func (self PileHeap) Len() int            { return len(self) }
func (self PileHeap) Less(i, j int) bool  { return self[i].Top() < self[j].Top() }
func (self PileHeap) Swap(i, j int)       { self[i], self[j] = self[j], self[i] }
func (self *PileHeap) Push(x interface{}) { *self = append(*self, x.(Pile)) }
func (self *PileHeap) Pop() interface{} {
	x := (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return x
}

func patienceSort(n []int) {
	var piles []Pile
	for _, x := range n {
		j := sort.Search(len(piles), func(i int) bool { return piles[i].Top() <= x })
		if j != len(piles) {
			piles[j] = append(piles[j], x)
		} else {
			piles = append(piles, Pile{x})
		}
	}

	ph := PileHeap(piles)
	heap.Init(&ph)
	for i, _ := range n {
		smallPile := heap.Pop(&ph).(Pile)
		n[i] = smallPile.Pop()
		if len(smallPile) != 0 {
			heap.Push(&ph, smallPile)
		}
	}
}

func main() {
	list := []int{91, 28, 73, 46, 55, 64, 37, 82, 19}
	fmt.Println("unsorted:", list)
	patienceSort(list)
	fmt.Println("sorted:  ", list)
}
