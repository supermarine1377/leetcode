package balanced_binary_tree

// Time Complexity: O(n) in case of full binary tree.
// Auxiliary Space: O(n) space for call stack since using recursionfunc isBalanced_ON(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, result := depthIsBalanced(root)
	return result
}

// Return the given tree's depth and whether the tree is balanced or not
func depthIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	ld, isBalancedL := depthIsBalanced(root.Left)
	rd, isBalancedR := depthIsBalanced(root.Right)

	if ld > rd {
		if ld-rd > 1 {
			return ld + 1, false
		}
		return ld + 1, isBalancedL && isBalancedR
	}
	if rd-ld > 1 {
		return rd + 1, false
	}
	return rd + 1, isBalancedL && isBalancedR
}
