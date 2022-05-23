package src

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val:%d Next:%+v}", n.Val, *n.Next)
}
