package tree

// pathSum 路径总和 II 返回路径
func pathSum(root *TreeNode, targetSum int) [][]int {

	if root == nil {
		return nil
	}

	// BFS
	result := make([][]int, 0)

	queue := make([]*PathSum, 0)
	queue = append(queue, &PathSum{root, root.Val})

	// 记录每个子树的父节点是什么， 当满足条件的时候遍历map 重新组成path
	parent := make(map[*TreeNode]*TreeNode)

	getPath := func(n *TreeNode) (path []int) {

		// 不断的向上找到父节点
		for ; n != nil; n = parent[n] {
			path = append(path, n.Val)
		}

		// path 原始是从低向上，现在反转
		for i, j := 0, len(path)-1; i < len(path)/2; i, j = i+1, j-1 {
			path[i], path[j] = path[j], path[i]
		}
		return
	}

	for len(queue) > 0 {

		for _, v := range queue {

			queue = queue[1:]

			if v.node.Left == nil && v.node.Right == nil && v.sum == targetSum {
				result = append(result, getPath(v.node))
			}

			if v.node.Left != nil {
				parent[v.node.Left] = v.node
				p := &PathSum{
					v.node.Left, v.sum + v.node.Left.Val,
				}
				queue = append(queue, p)
			}

			if v.node.Right != nil {
				parent[v.node.Right] = v.node
				p := &PathSum{v.node.Right, v.sum + v.node.Right.Val}
				queue = append(queue, p)
			}
		}
	}

	return result
}

type PathSum struct {
	// 当前节点
	node *TreeNode
	// 到当前节点node 是路径和
	sum int
	// 每次保存path 传递过去会导致内存溢出，可以用hmap 记录下来每个node 的parent
	// path []int
}

// pathSum 路径总和 II 返回路径
func pathSum2(root *TreeNode, targetSum int) [][]int {
	// DFS 从跟节点遍历，到也节点，剩余路径为0，则满足需求

	result := make([][]int, 0)

	path := make([]int, 0)

	var dfs func(*TreeNode, int)

	dfs = func(tn *TreeNode, left int) {
		if tn == nil {
			return
		}
		val := tn.Val
		left = targetSum - val
		// IMPORTANT
		// 为了防止path加入的节点一直存在，影响正常路径，需要 剪枝
		defer func() {
			path = path[:len(path)-1]
		}()
		if tn.Left == nil && tn.Right == nil && left == 0 {
			path = append(path, val)
			elem := make([]int, len(path))
			copy(elem, path)
			result = append(result, elem)
		}

		path = append(path, val)
		dfs(tn.Left, left)

		dfs(tn.Right, left)

	}

	dfs(root, targetSum)

	return result
}
