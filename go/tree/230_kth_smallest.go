package tree

import "sort"

// 因为二叉搜索树和中序遍历的性质，所以二叉搜索树的中序遍历是按照键增加的顺序进行的。
// 于是，我们可以通过中序遍历找到第 k 个最小元素。
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	// 中序遍历，迭代实现
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || node != nil {

		// inorder => 左子树，根节点，右子树
		// 先遍历左子树
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// 左子树遍历完，此时popped 为最左端没有左子树的节点了
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return popped.Val
		}

		// 右子树
		node = popped.Right
	}
	return 0
}

// kthSmallest 二叉搜索树 遍历后排序
func kthSmallest2(root *TreeNode, k int) int {

	if root == nil {
		return 0
	}

	// BFS 层序遍历
	data := make([]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		for _, v := range queue {
			queue = queue[1:]
			data = append(data, v.Val)

			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
	}

	sort.Ints(data)

	return data[k-1]
}
