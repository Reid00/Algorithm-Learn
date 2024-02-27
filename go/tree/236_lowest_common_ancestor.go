package tree

// lowestCommonAncestor 二叉树的最近公共祖先
// DFS
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 返回当前节点root
	// 公共祖先在当前节点以下的情况:
	// 1. 左右子树都有找到公共祖先
	// 2. 当前节点是nil
	// 3. 当前节点是p
	// 4. 当前节点是q

	// 其他
	// 公共祖先不在当前节点以下:
	// 5. 只有左子树找到，返回左子树递归的结果
	// 6. 只有右子树找到，返回右子树递归的结果
	// 7. 都没有找到，nil

	if root == nil {
		// 2
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		// 3,4
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		// 1
		return root
	}

	if left != nil {
		// 5
		return left
	}
	// 6
	return right

}

// 用hMap存储父节点
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	if root == nil {
		// 2
		return nil
	}

	parent := make(map[int]*TreeNode)

	var dfs func(*TreeNode)

	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}

		if tn.Left != nil {
			parent[tn.Left.Val] = tn
			dfs(tn.Left)
		}

		if tn.Right != nil {
			parent[tn.Right.Val] = tn
			dfs(tn.Right)
		}
	}

	dfs(root)
	visited := map[int]bool{}
	// 循环遍历获取p 的祖先
	for p != nil {
		// 用visited 记录p 的各个祖先
		visited[p.Val] = true
		p = parent[p.Val]
	}

	// 从低向上获取q 的祖先
	for q != nil {

		if visited[q.Val] {
			return q
		}

		q = parent[q.Val]
	}

	return nil
}
