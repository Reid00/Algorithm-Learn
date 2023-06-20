package tree

// invertTree 翻转二叉树  递归
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
