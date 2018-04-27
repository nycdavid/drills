package addtwonums

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// 321
	listAHead := NewReverseList(321)
	// 654
	listBHead := NewReverseList(654)

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
		t.Error("Expected value to be 5")
	}
}

func TestListToNum(t *testing.T) {
	listAHead := NewReverseList(321)
	number := listToNum(listAHead)

	if number != 321 {
		msg := fmt.Sprintf("Expected number to be 321, got %d", number)
		t.Error(msg)
	}
}

func TestLargeNumbers(t *testing.T) {
	op1 := 1234567
	op2 := 2345678
	listA := NewReverseList(op1)
	listB := NewReverseList(op2)
	resList := addTwoNumbers(listA, listB)
	result := listToNum(resList)

	expectedSum := op1 + op2
	if result != expectedSum {
		msg := fmt.Sprintf("Expected results to be %d, but got %d", expectedSum, result)
		t.Error(msg)
	}
}

func TestNegativeOperands(t *testing.T) {
	op1 := 10
	op2 := -5
	listA := NewReverseList(op1)
	listB := NewReverseList(op2)
	expectedSum := op1 + op2

	resList := addTwoNumbers(listA, listB)
	result := listToNum(resList)

	if result != expectedSum {
		msg := fmt.Sprintf("Expected results to be %d, but got %d", expectedSum, result)
		t.Error(msg)
	}
}
