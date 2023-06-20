package tree

// hasPathSum 路径总和  某个路径上是否存在到叶子节点总合等于targetSum 的路径
// BFS  加入到queue中的时候，同时把到此点的路径和计算出来
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	queue := make([]*Path, 0)
	queue = append(queue, &Path{root, root.Val})

	for len(queue) > 0 {

		for _, v := range queue {

			queue = queue[1:]

			if v.node.Left == nil && v.node.Right == nil && v.sum == targetSum {
				return true
			}

			if v.node.Left != nil {
				// v.node 的路径和 + 其左子树的val
				queue = append(queue, &Path{v.node.Left, v.sum + v.node.Left.Val})
			}

			if v.node.Right != nil {
				queue = append(queue, &Path{v.node.Right, v.sum + v.node.Right.Val})
			}
		}

	}
	return false
}

type Path struct {
	node *TreeNode
	sum  int
}

// hasPathSum 路径总和  某个路径上是否存在到叶子节点总合等于targetSum 的路径
// DFS
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum2(root.Left, targetSum-root.Val) ||
		hasPathSum2(root.Right, targetSum-root.Val)
}
