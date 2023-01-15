package maximum_depth_of_binary_tree

func maxDepthBFS_myans(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*TreeNode
	q = append(q, root)
	level := 0
	for len(q) > 0 {
		var newQ []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				newQ = append(newQ, node.Left)
			}
			if node.Right != nil {
				newQ = append(newQ, node.Right)
			}
		}
		q = newQ
		level++
	}
	return level
}
