package balanced_binary_tree

import "github.com/supermarine1377/leetcode/types"

type TreeNode = types.TreeNode

// Time Complexity: O(n^2) in case of full binary tree.
// Auxiliary Space: O(n) space for call stack since using recursion
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
