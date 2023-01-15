package average_of_levels_in_binary_tree

import "github.com/supermarine1377/leetcode/types"

type TreeNode = types.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		var nextQ []*TreeNode
		var tmpSum int
		div := len(q)
		for _, node := range q {
			tmpSum += node.Val
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}
		result = append(result, float64(tmpSum)/float64(div))
		q = nextQ
	}
	return result
}
