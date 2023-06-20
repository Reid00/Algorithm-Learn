package tree

import "math"

// widthOfBinaryTree二叉树最大宽度 BFS
// 假设当前节点的idx 是x 则，左子树是2x+1, 右子树是2x+2 (下标从0开始)
// 最大宽度就是每层的queue 中queue[last] - queue[0] + 1
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	width := math.MinInt32

	queue := make([]*Pair, 0)
	queue = append(queue, &Pair{0, root})

	for len(queue) > 0 {

		width = max(width, queue[len(queue)-1].idx-queue[0].idx+1)

		for _, v := range queue {
			// 此处循环的时候，有可能queue 里面通过下标queue[0], 或者queue[last] 获取的不仅仅是本层的元素
			// 虽然for range 遍历的queue 是固定的，但是通过下标获取的是最新的
			// width 的计算放到循环外面
			// width = max(width, queue[len(queue)-1].idx-queue[0].idx+1)
			queue = queue[1:]

			if v.node.Left != nil {
				// 此处的idx 应该为v的idx*2，而不是现有的queue 遍历时的i
				queue = append(queue, &Pair{2*v.idx + 1, v.node.Left})
			}

			if v.node.Right != nil {
				queue = append(queue, &Pair{2*v.idx + 2, v.node.Right})
			}

		}
	}

	return width
}

type Pair struct {
	idx  int
	node *TreeNode
}
