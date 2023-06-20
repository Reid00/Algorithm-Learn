package tree

// postorderTraversal 递归实现 (后序遍历)
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)

	if root == nil {
		return ans
	}

	var postorder func(*TreeNode)

	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		ans = append(ans, root.Val)

	}

	postorder(root)

	return ans
}

// postorderTraversal_ 迭代实现 (后序遍历)
func postorderTraversal_2(root *TreeNode) []int {

	ans := make([]int, 0)

	stack := make([]*TreeNode, 0)

	node := root

	// 引入了一个prev来记录历史访问记录
	var prev *TreeNode

	for len(stack) > 0 || node != nil {

		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// 走到此处说明所有的左孩子已经遍历到了，此时node == nil
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 后序遍历中，从栈中弹出的节点，我们只能确定其左子树肯定访问完了，但是无法确定右子树是否访问过。
		// 引入了一个prev来记录历史访问记录
		// 当访问完一棵子树的时候，我们用prev指向该节点。
		// 这样，在回溯到父节点的时候，我们可以依据prev是指向左子节点，还是右子节点，来判断父节点的访问情况。

		if popped.Right == nil || prev == popped.Right {
			//如果没有右子树， popped.Right == nil
			// 右子树访问完了，回到该节点的时候，需要判断该节点的右子树是否访问
			// 这个时候prev 用上了，如果prev == popped.Right 表示已经访问了改节点的右子树
			ans = append(ans, popped.Val)
			prev = popped
			popped = nil
		} else {
			// 否则有右孩子，右孩子应该先被访问，把popped (即该节点)压栈，再访问右子树
			stack = append(stack, popped)
			node = popped.Right
		}

	}

	return ans

}
