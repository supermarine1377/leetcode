package invert_biranary_tree

import "github.com/supermarine1377/leetcode/types"

type TreeNode = types.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
