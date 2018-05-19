package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeapProperty(t *testing.T) {
	arr := []int{10, 8, 12, 6, 9}
	maxHeap := NewMaxHeap(arr)

	expected := arr[2]
	got := maxHeap.Node(1).Val()

	if expected != got {
		msg := fmt.Sprintf("Expected %d, got %d", expected, got)
		t.Error(msg)
	}
}
