package addtwonums

import (
	"bytes"
	"strconv"
)

func addTwoNumbers(ln1 *ListNode, ln2 *ListNode) *ListNode {
	var newNode *ListNode
	var prevNode *ListNode
	var head *ListNode
	carry := 0

	for mrkr1, mrkr2 := ln1, ln2; continueLoop(mrkr1, mrkr2); mrkr1, mrkr2 = iterator(mrkr1, mrkr2) {
		num1 := assignVal(mrkr1)
		num2 := assignVal(mrkr2)

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

func iterator(mrkr1 *ListNode, mrkr2 *ListNode) (*ListNode, *ListNode) {
	if mrkr1 != nil {
		mrkr1 = mrkr1.Next
	}
	if mrkr2 != nil {
		mrkr2 = mrkr2.Next
	}
	return mrkr1, mrkr2
}

func continueLoop(mrkr1 *ListNode, mrkr2 *ListNode) bool {
	if mrkr1 == nil && mrkr2 == nil {
		return false
	}
	return true
}

func assignVal(mrkr *ListNode) int {
	if mrkr == nil {
		return 0
	} else {
		return mrkr.Val
	}
}
