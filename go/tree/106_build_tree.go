package tree

import "slices"

// buildTree_ 从中序与后序遍历序列构造二叉树
func buildTree_(inorder []int, postorder []int) *TreeNode {

	// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	// 输出：[3,9,20,null,null,15,7]

	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{rootVal, nil, nil}

	leftSize := slices.Index(inorder, rootVal)

	root.Left = buildTree_(inorder[:leftSize], postorder[:leftSize])
	root.Right = buildTree_(inorder[leftSize+1:], postorder[leftSize:len(postorder)-1])

	return root
}

// buildTree_2 从中序与后序遍历序列构造二叉树  迭代的方法实现
func buildTree_2(inorder []int, postorder []int) *TreeNode {

	// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
	// 输出：[3,9,20,null,null,15,7]

	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	var inorderIndex int = len(inorder) - 1

	for i := len(postorder) - 2; i >= 0; i-- {

		val := postorder[i]
		node := stack[len(stack)-1]

		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{val, nil, nil}
			stack = append(stack, node.Right)
		} else {

			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}

			node.Left = &TreeNode{val, nil, nil}
			stack = append(stack, node.Left)
		}

	}
	return root
}
