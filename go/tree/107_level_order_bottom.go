package tree

// levelOrderBottom 从低向上层次遍历
// 层序遍历后，反转
func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level := make([]int, 0)

		// 遍历当前这一层的结果
		// 也可以用 for 0..queue.len()
		for _, v := range queue {
			level = append(level, v.Val)
			queue = queue[1:]

			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}

		ans = append(ans, level)
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans

}
