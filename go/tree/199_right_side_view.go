package tree

func rightSideView(root *TreeNode) []int {
	// 层次遍历 队列先进先出
	// 正常从左到右入队列，现在从右到左

	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	ans := make([]int, 0)

	for len(queue) > 0 {
		// 第一个元素即为每层最右侧的元素
		ans = append(ans, queue[0].Val)

		for _, v := range queue {
			queue = queue[1:]

			if v.Right != nil {
				queue = append(queue, v.Right)
			}

			if v.Left != nil {
				queue = append(queue, v.Left)
			}
		}
	}
	return ans
}

// rightSideView  层序遍历最右侧的结果
func rightSideView2(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	ans := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		// level := make([]int, 0)
		// 优化
		ans = append(ans, queue[len(queue)-1].Val)

		for _, v := range queue {
			queue = queue[1:]
			// level = append(level, v.Val)

			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}

		}

		// ans = append(ans, level[len(level)-1])

	}

	return ans
}
