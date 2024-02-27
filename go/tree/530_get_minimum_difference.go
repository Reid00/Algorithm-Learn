package tree

import "math"

// getMinimumDifference 二叉搜索树的最小绝对差
// 中序遍历 递归的方式实现
func getMinimumDifference(root *TreeNode) int {
	// BST 中序遍历是从小到大的
	if root == nil {
		return 0
	}

	// 中序遍历 迭代的方式
	ans := make([]int, 0)
	var dfs func(*TreeNode)

	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}

		dfs(tn.Left)
		ans = append(ans, tn.Val)
		dfs(tn.Right)
	}

	dfs(root)

	var minVal = math.MaxInt32

	for i := 0; i < len(ans); i++ {

		if i > 0 {
			if ans[i]-ans[i-1] < minVal {
				minVal = ans[i] - ans[i-1]
			}
		}
	}

	return minVal
}

// getMinimumDifference 二叉搜索树的最小绝对差
// 中序遍历 迭代的方式
func getMinimumDifference2(root *TreeNode) int {
	// BST 中序遍历是从小到大的
	if root == nil {
		return 0
	}

	// 中序遍历 迭代的方式
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root

	for len(stack) > 0 || cur != nil {

		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		// 出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		cur = node.Right
	}

	var minVal = math.MaxInt32

	for i := 0; i < len(ans); i++ {

		if i > 0 {
			if ans[i]-ans[i-1] < minVal {
				minVal = ans[i] - ans[i-1]
			}
		}
	}

	return minVal
}
