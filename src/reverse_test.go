package src

import (
	"fmt"
	"testing"
)

var root *ListNode
var node2 *ListNode

func TestMain(t *testing.M) {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 = &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	root = &ListNode{Val: 0, Next: node1}
	fmt.Printf("%v\n", root)

	t.Run()
}

func TestReverse(t *testing.T) {
	fmt.Printf("%v\n", reverse(root))
}

func TestReverseRec(t *testing.T) {
	fmt.Printf("%v\n", reverseRec(root))
}

func TestReverseKGroup(t *testing.T) {
	fmt.Printf("%v\n", reverseKGroup(root, 2))
}

func TestReverseN(t *testing.T) {
	fmt.Printf("%v\n", reverseN(root, node2))
}
