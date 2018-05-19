package heap

import ()

type MaxHeap struct {
	heap *Heap
}

func NewMaxHeap(arr []int) *MaxHeap {
	heap := &Heap{arr: arr}
	return &MaxHeap{heap: heap}
}

func (*mh MaxHeap) Node(pos int) *HeapNode {
  
}
