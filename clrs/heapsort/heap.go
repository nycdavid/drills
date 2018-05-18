package heap

type Heap struct {
	arr []int
}

func (h *Heap) Length() int {
	return len(h.arr)
}
