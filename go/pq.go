package main

import (
	"container/heap"
)

type PriorityQueue struct {
	Data []interface{}
	Comparator func(a, b interface{}) int	
}

func NewPriorityQueue(comparator func(a, b interface{})int) *PriorityQueue {
	res := &PriorityQueue{
		Data: []interface{}{},
		Comparator: comparator,
	}
	heap.Init(res)
	return res
}

func (pq *PriorityQueue) Len() int {
	return len(pq.Data)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Comparator(pq.Data[i], pq.Data[j]) < 0
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Data[i], pq.Data[j] = pq.Data[j], pq.Data[i]
}

func (pq *PriorityQueue) Push(a interface{}) {
	pq.Data = append(pq.Data, a)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.Data)
	res := pq.Data[n-1]
	pq.Data = pq.Data[0:n-1]
	return res
}

func (pq *PriorityQueue) Put(a interface{}) {
	heap.Push(pq, a)
}

func (pq *PriorityQueue) Get() interface{} {
	return heap.Pop(pq)
}


