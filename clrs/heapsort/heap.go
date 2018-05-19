package heap

type Heap struct {
	arr []int
}

type HeapNode struct {
	arr []int
	Pos int
}

func (h *Heap) Length() int {
	return len(h.arr)
}

func (h *Heap) Node(pos int) *HeapNode {
	return &HeapNode{Pos: pos, arr: h.arr}
}

func (n *HeapNode) Left() *HeapNode {
	newPos := 2 * n.Pos
	return &HeapNode{Pos: newPos, arr: n.arr}
}

func (n *HeapNode) Right() *HeapNode {
	newPos := 2*n.Pos + 1
	return &HeapNode{Pos: newPos, arr: n.arr}
}

func (n *HeapNode) Parent() *HeapNode {
	newPos := n.Pos / 2
	return &HeapNode{Pos: newPos, arr: n.arr}
}

func (n *HeapNode) Val() int {
	idx := n.Pos - 1
	return n.arr[idx]
}
