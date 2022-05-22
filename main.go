package main

import (
	"fmt"

	. "github.com/ehlxr/leetcode-go/src"
)

func main() {
	node5 := &Node{Value: 5, Next: nil}
	node4 := &Node{Value: 4, Next: node5}
	node3 := &Node{Value: 3, Next: node4}
	node2 := &Node{Value: 2, Next: node3}
	node1 := &Node{Value: 1, Next: node2}
	root := &Node{Value: 0, Next: node1}

	// fmt.Printf("%v\n", root)

	// fmt.Printf("%v\n", Reverse(root))

	fmt.Printf("%v\n", ReverseKGroup(root, 2))
}
