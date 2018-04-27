package addtwonums

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// 321
	listATail := &ListNode{Val: 3, Next: nil}
	listAn1 := &ListNode{Val: 2, Next: listATail}
	listAHead := &ListNode{Val: 1, Next: listAn1}

	// 654
	listBTail := &ListNode{Val: 6, Next: nil}
	listBn1 := &ListNode{Val: 5, Next: listBTail}
	listBHead := &ListNode{Val: 4, Next: listBn1}

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
	listATail := &ListNode{Val: 3, Next: nil}
	listAn1 := &ListNode{Val: 2, Next: listATail}
	listAHead := &ListNode{Val: 1, Next: listAn1}

	number := listToNum(listAHead)

	if number != 321 {
		msg := fmt.Sprintf("Expected number to be 321, got %d", number)
		t.Error(msg)
	}
}
