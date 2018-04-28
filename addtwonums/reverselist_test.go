package addtwonums

import (
	"fmt"
	"testing"
)

func TestNewReverseList(t *testing.T) {
	operand := "54321"
	head := NewReverseList(operand)
	num := listToNum(head)

	if num != "54321" {
		msg := fmt.Sprintf("Expected num to be %s, got %s", "54321", num)
		t.Error(msg)
	}
}
