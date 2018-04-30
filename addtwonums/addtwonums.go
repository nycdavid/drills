package addtwonums

import (
	"bytes"
	"strconv"
)

type List struct {
	Head *ListNode
}

func addTwoNumbers(ln1 *ListNode, ln2 *ListNode) *ListNode {
	var newNode *ListNode
	var prevNode *ListNode
	var head *ListNode
	carry := 0

	mrkr2 := ln2
	for mrkr1 := ln1; mrkr1 != nil; mrkr1 = mrkr1.Next {
		num1 := mrkr1.Val
		num2 := mrkr2.Val

		newNum := num1 + num2 + carry
		if newNum > 9 {
			carry = 1
			newNum = newNum - 10
		} else {
			carry = 0
		}
		newNode = &ListNode{Val: newNum, Next: nil}

		if prevNode != nil {
			prevNode.Next = newNode
		} else {
			head = newNode
		}
		prevNode = newNode
		mrkr2 = mrkr2.Next
	}
	if carry != 0 {
		carry = 0
		newNode = &ListNode{Val: 1, Next: nil}
		prevNode.Next = newNode
	}
	return head
}

func listToNum(list *ListNode) string {
	var values []int
	for mrkr := list; mrkr != nil; mrkr = mrkr.Next {
		values = append(values, mrkr.Val)
	}
	var buf bytes.Buffer
	// values now hold the digits in reverse order
	for i := len(values); i != 0; i-- {
		idx := i - 1
		buf.WriteString(strconv.Itoa(values[idx]))
	}
	return buf.String()
}
