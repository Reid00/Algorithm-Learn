package tree

// inorderTraversal 中序遍历 （左根右），递归实现
func inorderTraversal(root *TreeNode) []int {

	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	var inorder func(*TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return ans
}

// inorderTraversal_ 中序遍历的迭代实现
func inorderTraversal_(root *TreeNode) []int {

	ans := make([]int, 0)

	stack := make([]*TreeNode, 0)

	node := root

	for len(stack) > 0 || node != nil {

		// 左根右顺序，所以要一直找左孩子
		// 先访问左子树
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 走出上面循环，node == nil， 左子树放完了，访问改节点，然后右子树
		// 最后一个左孩子即stack 中的最后一个元素
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans = append(ans, popped.Val)
		node = popped.Right
	}

	return ans
}