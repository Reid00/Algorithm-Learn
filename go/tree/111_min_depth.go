package tree

import "math"

// minDepth 二叉树最小深度 DFS
func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	ans := math.MaxInt32

	if root.Left != nil {
		ans = min(ans, minDepth(root.Left))
	}

	if root.Right != nil {
		ans = min(ans, minDepth(root.Right))
	}

	return ans + 1
}

// minDepth 二叉树最小深度 BFS
func minDepth2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := 1

	for len(queue) > 0 {

		for _, node := range queue {

			queue = queue[1:]

			// 遇到叶子结点，就是最小深度
			if node.Left == nil && node.Right == nil {
				return ans
			}

			if root.Left != nil {
				queue = append(queue, root.Left)
			}

			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		// 遍历完一层之后，ans  + 1
		ans += 1
	}
	return ans
}
