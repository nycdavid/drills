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
