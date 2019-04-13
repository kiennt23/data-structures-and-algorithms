package main

import (
	"container/heap"
	"fmt"
)

type IntegerHeap []int

func (iHeap IntegerHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntegerHeap) Less(i, j int) bool {
	return iHeap[i] < iHeap[j]
}

func (iHeap IntegerHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerHeap) Push(heapintf interface{}) {
	*iHeap = append(*iHeap, heapintf.(int))
}

func (iHeap *IntegerHeap) Pop() interface{} {
	previous := *iHeap
	n := len(previous)
	x1 := previous[n-1]
	*iHeap = previous[:n-1]
	return x1
}

func main() {
	intHeap := &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}

}
