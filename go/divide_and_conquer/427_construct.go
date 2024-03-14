package divideandconquer

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// 建立四叉树 construct
func construct(grid [][]int) *Node {
	// 1. 将格子四等分 => 横竖都切一半， 如果val 都相同为1/0 则是叶节点
	// 2. 如果某个四分之一格子 值不相同，则不是叶节点，将此格子四等分，观察是否可以是叶节点

	var dfs func(grid [][]int, colStart, colEnd int) *Node
	dfs = func(grid [][]int, colStart, colEnd int) *Node {

		for _, row := range grid {
			for _, v := range row[colStart:colEnd] {
				// 不是叶子节点
				if v != grid[0][colStart] {
					rMid, cMid := len(grid)/2, (colStart+colEnd)/2
					return &Node{
						Val:         false,
						IsLeaf:      false,
						TopLeft:     dfs(grid[:rMid], colStart, cMid),
						TopRight:    dfs(grid[:rMid], cMid, colEnd),
						BottomLeft:  dfs(grid[rMid:], colStart, cMid),
						BottomRight: dfs(grid[rMid:], cMid, colEnd),
					}
				}
			}
		}

		// 是叶子节点
		return &Node{Val: grid[0][colStart] == 1, IsLeaf: true}
	}

	return dfs(grid, 0, len(grid))
}
