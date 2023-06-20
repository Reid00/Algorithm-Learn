package tree

// preorderTraversal 前序遍历，递归实现
func preorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		ans = append(ans, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return ans
}

// preorderTraversal 前序遍历，迭代实现
func preorderTraversal_(root *TreeNode) []int {

	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	// 栈存储，根，左，右
	stack := make([]*TreeNode, 0)

	node := root

	for len(stack) > 0 || node != nil {

		for node != nil {
			ans = append(ans, node.Val)
			// 将根节点 压栈 方便后续取出其右孩子
			stack = append(stack, node)
			node = node.Left
		}
		// 压栈完成，根节点全部放到ans 中
		// 出栈获取 右孩子
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = popped.Right
	}

	return ans
}
