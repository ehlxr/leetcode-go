package src

func reverse(root *ListNode) *ListNode {
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

/* 递归反转链表 */
func reverseRec(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	res := reverseRec(root.Next)
	root.Next.Next = root
	root.Next = nil

	return res
}

/* K 个一组翻转链表 */
func reverseKGroup(root *ListNode, k int) *ListNode {
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
	tmp := reverseN(root, n)
	root.Next = reverseKGroup(n, k)
	return tmp
}

/* 反转链表到指定节点 */
func reverseN(root *ListNode, n *ListNode) *ListNode {
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
