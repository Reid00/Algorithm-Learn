package tree

// sumNumbers 求根节点到叶节点数字之和
// METHOD 1: DFS
func sumNumbers(root *TreeNode) int {
	return AddPrev(root, 0)
}

func AddPrev(root *TreeNode, prev int) int {
	if root == nil {
		return 0
	}
	tmp := prev*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return tmp
	}

	return AddPrev(root.Left, tmp) + AddPrev(root.Right, tmp)
}

// BFS， 每遍历一层，记录当前层的TreeNode 和 累积和
func sumNumbers2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var result int

	queue := make([]Sum, 0)
	queue = append(queue, Sum{root, 0})

	for len(queue) > 0 {

		for _, v := range queue {
			queue = queue[1:]

			tmp := v.sum*10 + v.node.Val

			if v.node.Left == nil && v.node.Right == nil {
				// 每层叶子节点的结果 加到result 上
				result += tmp
			}

			if v.node.Left != nil {
				queue = append(queue, Sum{v.node.Left, tmp})
			}

			if v.node.Right != nil {
				queue = append(queue, Sum{v.node.Right, tmp})
			}
		}
	}

	return result
}

type Sum struct {
	node *TreeNode
	// sum: 到 node 节点 父节点的累积和
	sum int
}
