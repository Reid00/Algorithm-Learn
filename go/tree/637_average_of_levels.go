package tree

// averageOfLevels 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。
// BFS 每层的结果，然后计算值
func averageOfLevels(root *TreeNode) []float64 {

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
