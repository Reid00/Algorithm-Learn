package tree

// BFS levelOrder 层级遍历，queue
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		// 当前层的节点个数
		length := len(queue)
		level := make([]int, 0)
		// also can use for range instead
		for i := 0; i < length; i++ {
			head := queue[0]
			level = append(level, head.Val)
			queue = queue[1:]

			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}

		ans = append(ans, level)
	}

	return ans

}
