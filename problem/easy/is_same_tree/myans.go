package is_same_tree

import (
	"github.com/supermarine1377/leetcode/types"
)

type TreeNode = types.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	var pQ []*TreeNode
	pQ = append(pQ, p)
	var qQ []*TreeNode
	qQ = append(qQ, q)
	for len(pQ) != 0 && len(qQ) != 0 {
		if !isEqual(pQ, qQ) {
			return false
		}
		pQ = newQue(pQ)
		qQ = newQue(qQ)
	}
	return true
}

func newQue(q []*TreeNode) []*TreeNode {
	var newQ []*TreeNode
	for _, n := range q {
		if n != nil {
			newQ = append(newQ, n.Left)
			newQ = append(newQ, n.Right)
		}
	}
	return newQ
}

func isEqual(p []*TreeNode, q []*TreeNode) bool {
	for i := range p {
		if p[i] == nil && q[i] != nil {
			return false
		}
		if p[i] != nil && q[i] == nil {
			return false
		}
		if p[i] == nil && q[i] == nil {
			continue
		}
		if p[i].Val != q[i].Val {
			return false
		}
	}
	return true
}

func isSameTreeRecursively(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
