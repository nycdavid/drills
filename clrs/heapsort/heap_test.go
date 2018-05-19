package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapArr(t *testing.T) {
	heap := &Heap{}
	typ := reflect.TypeOf(heap.arr)

	got := typ.Kind().String()
	expected := "slice"

	if got != expected {
		msg := fmt.Sprintf("Expected kind to be %s, got %s", expected, got)
		t.Error(msg)
	}
}

func TestHeapLength(t *testing.T) {
	heap := &Heap{}

	got := heap.Length()
	expected := 0

	if got != expected {
		msg := fmt.Sprintf("Expected Length to be %d, got %d", expected, got)
		t.Error(msg)
	}
}

func TestHeapNode(t *testing.T) {
	heapArr := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	heap := &Heap{arr: heapArr}

	expected := 16
	got := heap.Node(1).Val()

	if expected != got {
		msg := fmt.Sprintf("Expected heapNode to be %d, got %d", expected, got)
		t.Error(msg)
	}
}

func TestHeapNodeLeft(t *testing.T) {
	heapArr := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	heap := &Heap{arr: heapArr}
	node := heap.Node(1)

	expected := heapArr[1]
	got := node.Left().Val()

	if expected != got {
		msg := fmt.Sprintf("Expected left to be %d, got %d", expected, got)
		t.Error(msg)
	}
}

func TestHeapNodeRight(t *testing.T) {
	heapArr := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	heap := &Heap{arr: heapArr}
	node := heap.Node(1)

	expected := heapArr[2]
	got := node.Right().Val()

	if expected != got {
		msg := fmt.Sprintf("Expected left to be %d, got %d", expected, got)
		t.Error(msg)
	}
}

func TestHeapNodeParent(t *testing.T) {
	heapArr := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	heap := &Heap{arr: heapArr}
	node := heap.Node(7)

	expected := heapArr[2]
	got := node.Parent().Val()

	if expected != got {
		msg := fmt.Sprintf("Expected left to be %d, got %d", expected, got)
		t.Error(msg)
	}
}
