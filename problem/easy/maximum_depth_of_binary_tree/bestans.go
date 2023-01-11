package maximum_depth_of_binary_tree

import "github.com/supermarine1377/leetcode/types"

type TreeNode = types.TreeNode

func maxDepthRecursively(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := maxDepthRecursively(root.Left), maxDepthRecursively(root.Right)
	if leftDepth > rightDepth {
		return 1 + leftDepth
	}
	return 1 + rightDepth
}
