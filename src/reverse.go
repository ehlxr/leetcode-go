package src

func Reverse(root *Node) *Node {
	var pre, tmp *Node
	cur := root
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

	}
	return pre
}

func ReverseKGroup(root *Node, k int) *Node {
	if root == nil {
		return root
	}
	n := root
	for i := 0; i < k; i++ {
		if n == nil {
			break
		}
		n = n.Next
	}
	tmp := ReverseN(root, n)
	root.Next = ReverseKGroup(n, k)
	return tmp
}

func ReverseN(root *Node, n *Node) *Node {
	var pre, tmp *Node
	cur := root
	for cur != n {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

	}
	return pre
}
