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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (n *TreeNode) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("{Val:%d Left:%+v Right:%+v}", n.Val, *n.Left, *n.Right)
}
