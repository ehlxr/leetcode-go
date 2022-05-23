package src

func Reverse(root *ListNode) *ListNode {
	var pre, tmp *ListNode
	cur := root
	for cur != nil {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

	}
	return pre
}

func ReverseRec(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	res := ReverseRec(root.Next)
	root.Next.Next = root
	root.Next = nil

	return res
}

/* K 个一组翻转链表 */
func ReverseKGroup(root *ListNode, k int) *ListNode {
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

/* 反转链表到指定节点 */
func ReverseN(root *ListNode, n *ListNode) *ListNode {
	var pre, tmp *ListNode
	cur := root
	for cur != n {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp

	}
	return pre
}
