package addtwonums

import (
	"bytes"
	"log"
	"strconv"
)

type List struct {
	Head *ListNode
}

func addTwoNumbers(ln1 *ListNode, ln2 *ListNode) *ListNode {
	operand1 := listToNum(ln1)
	operand2 := listToNum(ln2)
	sum := operand1 + operand2
	strSum := strconv.Itoa(sum)

	num, _ := strconv.Atoi(string(strSum[0]))
	node := &ListNode{Val: num, Next: nil}
	list := &List{Head: node}

	for i := 1; i < len(strSum); i++ {
		num, _ := strconv.Atoi(string(strSum[i]))
		headNode := list.Head
		node = &ListNode{Val: num, Next: headNode}
		list.Head = node
	}

	return list.Head
}

func listToNum(list *ListNode) int {
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
	num, err := strconv.Atoi(buf.String())

	if err != nil {
		log.Print(err)
		panic(err)
	}
	return num
}
