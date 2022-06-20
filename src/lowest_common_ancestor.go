package src

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先
// 对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return find(root, p.Val, q.Val)
}

func find(root *TreeNode, l, r int) *TreeNode {
	if root == nil {
		return nil
	}
	// p 和 q 均存在于给定的二叉树中
	if root.Val == l || root.Val == r {
		return root
	}

	left := find(root.Left, l, r)
	right := find(root.Right, l, r)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	// p 和 q 均存在于给定的二叉树中
// 	if root.Val == p.Val || root.Val == q.Val {
// 		return root
// 	}
//
// 	left := lowestCommonAncestor(root.Left, p, q)
// 	right := lowestCommonAncestor(root.Right, p, q)
// 	if left != nil && right != nil {
// 		return root
// 	}
//
// 	if left != nil {
// 		return left
// 	}
// 	return right
// }
