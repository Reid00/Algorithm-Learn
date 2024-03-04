package graph

// islandPerimeter 岛屿周长
func islandPerimeter(grid [][]int) int {

	// 两种情况贡献周长
	// 1. 格子的周边是边缘，下一步dfs 超出格子范围时，贡献1
	// 2. 格子是岛屿，周边是海域的情况时，贡献1

	var res int

	var dfs func([][]int, int, int) int
	dfs = func(grid [][]int, r, c int) int {

		// 超出边界贡献一条
		if !inArea(grid, r, c) {
			return 1
		}

		// 处于海域
		if grid[r][c] == 0 {
			return 1
		}

		// 因为「当前格子是已遍历的陆地格子」返回，和周长没关系
		if grid[r][c] != 1 {
			return 0
		}

		// 标记已经访问过
		grid[r][c] = 2

		return dfs(grid, r+1, c) +
			dfs(grid, r-1, c) +
			dfs(grid, r, c+1) +
			dfs(grid, r, c-1)
	}

	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				ret := dfs(grid, i, j)
				res = max(res, ret)
			}
		}
	}

	return res
}
