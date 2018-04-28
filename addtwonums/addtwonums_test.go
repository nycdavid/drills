package addtwonums

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbersOfEqualLength(t *testing.T) {
	// 321
	listAHead := NewReverseList("321")
	// 654
	listBHead := NewReverseList("654")

	resList := addTwoNumbers(listAHead, listBHead)
	var values []int
	// answer is 975
	for mrkr := resList; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}

	if len(values) != 3 {
		msg := fmt.Sprintf("Expected 3 elements, got %d", len(values))
		t.Error(msg)
	}
	if resList.Val != 5 {
		msg := fmt.Sprintf("Expected value to be %d, got %d", 5, resList.Val)
		t.Error(msg)
	}
}

func TestAddTwoNumbersOfDifferingLength(t *testing.T) {
	// 2334
	// listAHead := NewReverseList("2334")
	// 125
	// listBHead := NewReverseList("125")
}

func TestListToNum(t *testing.T) {
	listAHead := NewReverseList("321")
	number := listToNum(listAHead)

	if number != "321" {
		msg := fmt.Sprintf("Expected number to be 321, got %s", number)
		t.Error(msg)
	}
}

func TestLargeNumbers(t *testing.T) {
	op1 := "1234567"
	op2 := "2345678"
	listA := NewReverseList(op1)
	listB := NewReverseList(op2)
	resList := addTwoNumbers(listA, listB)
	result := listToNum(resList)

	if result != "3580245" {
		msg := fmt.Sprintf("Expected results to be %s, but got %s", "3580245", result)
		t.Error(msg)
	}
}
