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

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		for _, node := range q {
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return level
}

func maxDepthDFS(root *TreeNode) int {
	type node struct {
		*TreeNode
		level int
	}
	var st []node
	st = append(st, node{root, 1})
	result := 0
	for len(st) != 0 {
		bottom := st[len(st)-1]
		st = st[0 : len(st)-1]
		if bottom.TreeNode != nil {
			if bottom.level > result {
				result = bottom.level
			}
			st = append(st, node{bottom.Left, bottom.level + 1})
			st = append(st, node{bottom.Right, bottom.level + 1})
		}
	}
	return result
}
