package tree

// connect 填充每个节点的下一个右侧节点指针
// DFS 递归的方式
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	var connectTwoNodes func(*Node, *Node)

	connectTwoNodes = func(n1, n2 *Node) {

		if n1 == nil || n2 == nil {
			return
		}

		n1.Next = n2

		// 有相同的父节点子节点指定next
		connectTwoNodes(n1.Left, n1.Right)
		connectTwoNodes(n2.Left, n2.Right)

		// 不同父节点的子节点指定
		connectTwoNodes(n1.Right, n2.Left)

	}
	connectTwoNodes(root.Left, root.Right)
	return root
}

// connect 填充每个节点的下一个右侧节点指针
// BFS 每一层的右侧元素是 queue index 的下一个
func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		// 遍历每一层
		length := len(queue)

		for i := 0; i < length; i++ {
			// 因为n 始终是queue[0] next 就是queue[1]
			n := queue[0]
			if i+1 < length {
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

		// queue = queue[length:]
	}

	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
