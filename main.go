package main

import "fmt"

func main() {
	// node5 := &ListNode{Val: 5, Next: nil}
	// node4 := &ListNode{Val: 4, Next: node5}
	// node3 := &ListNode{Val: 3, Next: node4}
	// node2 := &ListNode{Val: 2, Next: node3}
	// node1 := &ListNode{Val: 1, Next: node2}
	// root := &ListNode{Val: 0, Next: node1}
	// fmt.Printf("%v\n", root)
	//
	// fmt.Printf("%v\n", Reverse(root))
	// fmt.Printf("%v\n", ReverseRec(root))
	// fmt.Printf("%v\n", ReverseKGroup(root, 2))

	s := "Hello 你好"
	fmt.Println(s[2], s[6])
	rs := []rune(s)
	println(string(rs[6]))

	// s := "babad"
	// println(LongestPalindrome(s))
}
