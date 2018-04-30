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
	for idx, digit := range []int{5, 7, 9} {
		if values[idx] != digit {
			msg := fmt.Sprintf("Expected el at idx %d to equal %d, but got %d", idx, digit, values[idx])
			t.Error(msg)
		}
	}
}

func TestAddTwoNumbersOfDifferingLength(t *testing.T) {
	listAHead := NewReverseList("2334")
	listBHead := NewReverseList("125")
	resList := addTwoNumbers(listAHead, listBHead)
	var values []int
	for mrkr := resList; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}

	if len(values) != 4 {
		msg := fmt.Sprintf("Expected 4 elements, got %d", len(values))
		t.Error(msg)
	}
	for idx, digit := range []int{9, 5, 4, 2} {
		if values[idx] != digit {
			msg := fmt.Sprintf("Expected el at idx %d to equal %d, but got %d", idx, digit, values[idx])
			t.Error(msg)
		}
	}
}

func TestHandlesCarryOver(t *testing.T) {
	listAHead := NewReverseList("9")
	listBHead := NewReverseList("9")
	resList := addTwoNumbers(listAHead, listBHead)
	var values []int
	for mrkr := resList; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}

	if len(values) != 2 {
		msg := fmt.Sprintf("Expected 2 elements, got %d", len(values))
		t.Error(msg)
	}
	for idx, digit := range []int{8, 1} {
		if values[idx] != digit {
			msg := fmt.Sprintf("Expected el at idx %d to equal %d, but got %d", idx, digit, values[idx])
			t.Error(msg)
		}
	}
}

func TestHandledMultipleCarryOver(t *testing.T) {
	listAHead := NewReverseList("99")
	listBHead := NewReverseList("99")
	resList := addTwoNumbers(listAHead, listBHead)
	var values []int
	for mrkr := resList; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}

	expectedLen := 3
	if len(values) != expectedLen {
		msg := fmt.Sprintf("Expected %d elements, got %d", expectedLen, len(values))
		t.Error(msg)
	}
	for idx, digit := range []int{8, 9, 1} {
		if values[idx] != digit {
			msg := fmt.Sprintf("Expected el at idx %d to equal %d, but got %d", idx, digit, values[idx])
			t.Error(msg)
		}
	}
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
	listA := NewReverseList("1234567")
	listB := NewReverseList("2345678")
	resList := addTwoNumbers(listA, listB)
	result := listToNum(resList)
	var values []int
	for mrkr := resList; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}

	if result != "3580245" {
		msg := fmt.Sprintf("Expected results to be %s, but got %s", "3580245", result)
		t.Error(msg)
	}
	for idx, digit := range []int{5, 4, 2, 0, 8, 5, 3} {
		if values[idx] != digit {
			msg := fmt.Sprintf("Expected el at idx %d to equal %d, but got %d", idx, digit, values[idx])
			t.Error(msg)
		}
	}
}
