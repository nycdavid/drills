package addtwonums

import (
	"log"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewReverseList(operand string) *ListNode {
	node := createNode(operand[0])
	var prevNode *ListNode
	for i := 1; i < len(operand); i++ {
		prevNode = node
		node = createNode(operand[i])
		node.Next = prevNode
	}
	return node
}

func createNode(chr byte) *ListNode {
	num, err := strconv.Atoi(string(chr))
	if err != nil {
		log.Print(err)
	}
	return &ListNode{Val: num, Next: nil}
}
