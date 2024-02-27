package tree

func averageOfLevels(root *TreeNode) []float64 {
	// BFS

	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	ans := make([]float64, 0)

	for len(queue) > 0 {

		// 计算当前层的平均值
		var sum float64
		length := len(queue)

		for _, v := range queue {
			sum += float64(v.Val)
			queue = queue[1:]

			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}

		avg := sum / float64(length)

		ans = append(ans, avg)

	}
	return ans
}

// averageOfLevels 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。
// BFS 每层的结果，然后计算值
func averageOfLevels2(root *TreeNode) []float64 {

	if root == nil {
		return nil
	}

	ans := make([]float64, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var level float64

		cnt := len(queue)

		for i := 0; i < cnt; i++ {
			head := queue[0]
			queue = queue[1:]

			level += float64(head.Val)

			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}

		avg := level / float64(cnt)

		ans = append(ans, avg)
	}

	return ans

}
