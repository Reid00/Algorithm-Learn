package tree

// isSymmetric 对称二叉树 BFS 实现
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	// 遍历每一层的数据，查看是否相同
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		length := len(queue)
		// 此处判断对称逻辑不完整, 
        // 加上node.Left == nil 时append nil, 同时判断queue[i] 是否为nil
		for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
            if queue[i] == nil && queue[j] == nil {
                continue
            }
            if queue[i] == nil || queue[j] == nil {
                return false
            }

			if queue[i].Val != queue[j].Val {
				return false
			}
		}

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				continue
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			} else {
				queue = append(queue, nil)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			} else {
				queue = append(queue, nil)
			}
		}

	}
	return true
}

// isSymmetric 对称二叉树 DFS 实现
func isSymmetric2(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var check func(*TreeNode, *TreeNode) bool

	check = func(left, right *TreeNode) bool {
		// 都为nil
		if left == nil && right == nil {
			return true
		}
		// 一个为nil 另一个补位nil false
		if left == nil || right == nil {
			return false
		}

		// 值相同，然后递归比较左右子树 是否对称
		return left.Val == right.Val && check(left.Left, right.Right) &&
			check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}
