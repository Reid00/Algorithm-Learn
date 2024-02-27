package tree

// connect 填充每个节点的下一个右侧节点指针, 非完美二叉树
// O(n) 空间复杂度
func connect3(root *Node) *Node {

	// BFS
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		length := len(queue)

		for i := 0; i < length; i++ {

			n := queue[0]
			if i < length-1 {
				n.Next = queue[1]
			} else {
				n.Next = nil
			}

			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}

		}

	}

	return root
}

// O(1) 空间复杂度, BFS + 链表
func connect_(root *Node) *Node {

	// BFS
	if root == nil {
		return nil
	}

	dummy := &Node{}
	cur := root

	for cur != nil {

		dummy.Next = nil // 哨兵节点，dummy.Next 是下一层的起始节点, NOTE: 要置为空，让外层退出循环
		prev := dummy

		// 在i层时， 为i+1 层建立next 关系， 这样dummy 操作的就是指向下一层了
		for cur != nil { // 遍历当前层

			if cur.Left != nil {
				// 第一次prev.Next = cur.Left 也会让dummy.Next 指向下一层的最左侧节点
				prev.Next = cur.Left
				prev = cur.Left
			}

			if cur.Right != nil {
				prev.Next = cur.Right
				prev = cur.Right
			}

			// 当前层的下一个节点
			cur = cur.Next
		}

		// 下一层的头节点
		cur = dummy.Next
	}

	return root
}
