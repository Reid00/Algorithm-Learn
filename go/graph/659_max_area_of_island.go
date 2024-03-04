package graph

// maxAreaOfIsland 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	// DFS
	if len(grid) == 0 {
		return 0
	}

	var res = 0

	row, col := len(grid), len(grid[0])

	var dfs func([][]int, int, int) int
	dfs = func(grid [][]int, i2, i3 int) int {

		// base case
		if !inArea(grid, i2, i3) {
			return 0
		}

		// 如果这个格子不是岛屿，直接返回
		if grid[i2][i3] != 1 {
			return 0
		}

		// 是岛屿标记为访问过的
		grid[i2][i3] = 2

		// 返回面积，是岛屿 1 + 邻居
		return 1 +
			dfs(grid, i2+1, i3) +
			dfs(grid, i2-1, i3) +
			dfs(grid, i2, i3+1) +
			dfs(grid, i2, i3-1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, i, j)
				res = max(res, area)
			}
		}
	}
	return res
}
