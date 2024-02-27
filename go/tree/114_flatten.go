package tree

// flatten 二叉树展开为链表 效果和先序遍历 顺序相同
// O(1)
func flatten(root *TreeNode) {
	// 1. 寻找当前节点左子树最右边的节点作为前驱节点(predecessor)
	// 2. 讲当前节点的右子树作为predecessor 的右子树
	// 3. 讲当前节点的左子树(next)更新为右子树，左子树置空
	// 4. 循环cur = cur.Left

	cur := root

	for cur != nil {

		if cur.Left != nil {

			next := cur.Left
			pre := next

			// 找到当前左子树左右侧的节点
			for pre.Right != nil {
				pre = pre.Right
			}

			// 将当前节点的右子树，挂到 pre 的右子树上
			pre.Right = cur.Right
			cur.Left, cur.Right = nil, next
		}
		cur = cur.Right
	}
}

// flatten 二叉树展开为链表 效果和先序遍历 顺序相同
// O(n)
func flatten_(root *TreeNode) {
	// METHOD 1: 前序遍历存储到切片中， 然后遍历切片修改Left=nil, Right=i+1的元素
	if root == nil {
		return
	}

	// 前序层次遍历， 根左右
	stack := make([]*TreeNode, 0)
	node := root

	result := make([]*TreeNode, 0)
	// preorder
	for len(stack) > 0 || node != nil {

		for node != nil {
			stack = append(stack, node)
			result = append(result, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]

	}
	// 前序遍历
	for i, n := range result {
		// 最后一个
		if i == len(result)-1 {
			n.Left = nil
			n.Right = nil
		} else {
			n.Left = nil
			n.Right = result[i+1]

		}
	}

}
