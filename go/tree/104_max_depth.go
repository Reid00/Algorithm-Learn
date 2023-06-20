package tree

// maxDepth DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// maxDepth BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		for _, v := range queue {

			queue = queue[1:]
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		ans += 1
	}

	return ans

}
