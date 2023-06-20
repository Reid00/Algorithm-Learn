package tree

// buildTree 从前序与中序遍历序列构造二叉树
// 前序 跟左右  => 第一个是root
// 中序 左根右  => 根据上面的root确定其左边是左子树
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}
	// root 在中序遍历的下标
	rootIdxInorder := findIdx(inorder, rootVal)

	// 左子树的元素个数为: rootIdxInorder (下标从0开始)
	root.Left = buildTree(preorder[1:1+rootIdxInorder], inorder[:rootIdxInorder])
	root.Right = buildTree(preorder[rootIdxInorder+1:], inorder[rootIdxInorder+1:])
	return root
}

func findIdx(nums []int, val int) int {

	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return -1
}
