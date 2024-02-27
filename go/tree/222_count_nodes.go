package tree

import "math"

// countNode 完全二叉树的节点个数
func countNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var cnt int

	// BFS
	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {

		for _, v := range q {
			cnt += 1
			q = q[1:]

			if v.Left != nil {
				q = append(q, v.Left)
			}

			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
	}
	return cnt
}

// countNode 完全二叉树的节点个数
// DFS
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 计算深度
	lDepth := 0
	rDepth := 0

	// 左子树深度
	for node := root; node != nil; node = node.Left {
		lDepth += 1
	}
	// 右
	for node := root; node != nil; node = node.Right {
		lDepth += 1
	}

	if lDepth == rDepth {
		// return int(math.Pow(2, float64(lDepth)+1) - 1)
		return 2 << (lDepth+1) -1
	}

	return countNodes2(root.Left) + countNodes2(root.Right) + 1
}
