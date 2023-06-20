package tree

import "math"

// isValidBST 是否是一个有效的二叉搜索树
// 二叉搜索树 左子树都比根节点小，右子树都比根节点大
// DFS
func isValidBST(root *TreeNode) bool {

	var helper func(*TreeNode, int, int) bool

	helper = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}

		// lower, upper 代表左子树和右子树的值
		// BST 根节点应该大于左子树的值，小于右子树的值
		if node.Val <= lower || node.Val >= upper {
			return false
		}

		return helper(node.Left, lower, node.Val) && helper(node.Right, node.Val, upper)
	}

	return helper(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历
func isValidBST2(root *TreeNode) bool {

	if root == nil {
		return true
	}

	// 中序遍历，迭代实现
	stack := make([]*TreeNode, 0)

	// ans := make([]int, 0)
	prev := math.MinInt64

	node := root

	for len(stack) > 0 || node != nil {

		// 先遍历左子树
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 根节点
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// ans = append(ans, popped.Val)
		// if len(ans) > 1 && ans[len(ans)-1] <= ans[len(ans)-2] {
		// 	return false
		// }
		// 优化上面的ans ，不用O(n) 的空间
		if popped.Val <= prev {
			return false
		}
		prev = popped.Val
		// if popped.Val <= node.Val {
		// 	return false
		// }
		// 右子树
		node = popped.Right
	}
	return true
}
