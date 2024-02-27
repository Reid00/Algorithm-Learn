package tree

import (
	"slices"
)

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
	leftSize := findIdx(inorder, rootVal)

	// 左子树的元素个数为: leftSize (下标从0开始)
	root.Left = buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	root.Right = buildTree(preorder[leftSize+1:], inorder[leftSize+1:])
	return root
}

func findIdx(nums []int, val int) int {

	// for i, v := range nums {
	// 	if v == val {
	// 		return i
	// 	}
	// }
	// return -1
	return slices.Index(nums, val)
}

// 迭代的方法实现
func buildTree2(preorder []int, inorder []int) *TreeNode {

	// preorder = [ 3, 9, 20, 15, 7 ]
	// inorder = [ 20, 9, 15, 3, 7 ]
	// 输出: [3,9,7,20,15,null,null]

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	// 栈模拟递归
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	var inorderIndex int

	// 把跟节点以外的preorder 入栈 前序 根左右
	for i := 1; i < len(preorder); i++ {

		val := preorder[i]
		node := stack[len(stack)-1]

		// 栈顶元素 和 inorder 值不相同，继续是左子树入栈
		if node.Val != inorder[inorderIndex] {

			node.Left = &TreeNode{val, nil, nil}
			stack = append(stack, node.Left)
		} else {

			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				// 出栈
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				inorderIndex++
			}

			node.Right = &TreeNode{Val: val}
			stack = append(stack, node.Right)
		}
	}

	return root
}
