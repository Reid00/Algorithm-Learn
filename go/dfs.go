package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历递归实现
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return result

}

// 前序遍历非递归实现
func preorderTraversalNo(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	// stack = append(stack, root) #BUG will be if add at first
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			result = append(result, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		// node = stack[len(stack)-1] # BUG should be get the right node
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

	}
	return result
}

// 中序遍历递归
func indorderTraversal(root *TreeNode) []int {
	var result []int

	var inorder func(*TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return result
}

// 中序遍历非递归
func indorderTraversalNo(root *TreeNode) []int {
	var result []int

	var stack []*TreeNode
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		node = val.Right

	}
	return result
}

// 后序遍历递归
func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var postorder func(*TreeNode)

	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		result = append(result, root.Val)
	}

	postorder(root)
	return result
}

// 后序遍历非递归实现
func postorderTraversalNo(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	// 前序遍历 根 左右
	// 后续遍历 [左右根] 先记录 根右左  然后倒序-> 左右根
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			result = append(result, node.Val)
			node = node.Right
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Left
	}
	// 倒序result
	res := make([]int, 0, len(result))
	for i := range result {
		res = append(res, result[len(result)-1-i])
	}
	return res
}
