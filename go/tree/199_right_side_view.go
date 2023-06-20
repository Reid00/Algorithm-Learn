package tree

// rightSideView  层序遍历最右侧的结果
func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	ans := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		level := make([]int, 0)

		for _, v := range queue {
			queue = queue[1:]
			level = append(level, v.Val)

			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}

		}

		ans = append(ans, level[len(level)-1])

	}

	return ans

}
