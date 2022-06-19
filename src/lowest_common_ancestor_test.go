package src

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	// 初始化二叉树
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 4}

	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 5, 1, lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 1}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 5, 4, lowestCommonAncestor(root, &TreeNode{Val: 5}, &TreeNode{Val: 4}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 2, 6, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 6}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 2, 7, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 7}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 2, 4, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 4}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 2, 3, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 3}).Val)
	fmt.Printf("lowestCommonAncestor of %d and %d is %+v \n", 2, 1, lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 1}).Val)
}
