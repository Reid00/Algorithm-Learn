package tree

// isSameTree 相同的树
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	// 优化
	if p == nil || q == nil {
		return false
	}

	// if p != nil && q == nil || p == nil && q != nil {
	// 	return false
	// }

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}