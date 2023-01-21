package balanced_binary_tree

import "github.com/supermarine1377/leetcode/types"

type TreeNode = types.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ld, rd := depth(root.Left), depth(root.Right)
	if ld-rd > 1 || rd-ld > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld, rd := depth(root.Left), depth(root.Right)
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}
