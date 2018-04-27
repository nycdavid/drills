package addtwonums

import (
	"testing"
)

func TestNewReverseList(t *testing.T) {
	operand := 54321
	head := NewReverseList(operand)
	num := listToNum(head)

	if num != 54321 {
		t.Error("Expected num to be %d, got %d", 54321, num)
	}
}
