package tree

//  * Definition for a Node.
type MNode struct {
	Val      int
	Children []*MNode
}

// BFS N 叉树的最大深度
func maxDepth_N(root *MNode) int {

	if root == nil {
		return 0
	}

	ans := 0

	queue := make([]*MNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		for _, v := range queue {
			queue = queue[1:]
			// 遍历v 的孩子，加入queue 中
			if v.Children == nil {
				break
			}

			for _, node := range v.Children {
				queue = append(queue, node)
			}

		}
		ans += 1
	}
	return ans
}

// DFS N 叉树的最大深度
func maxDepth_N_2(root *MNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0

	for _, v := range root.Children {
		if childDepth := maxDepth_N_2(v); childDepth > maxDepth {
			maxDepth = childDepth
		}
	}

	return maxDepth + 1
}
