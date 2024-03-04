package graph

// numIslands 岛屿数量
func numIslands(grid [][]byte) int {
	// BFS
	// 碰到1 的岛屿加入队列中，然后出队列

	if len(grid) == 0 {
		return 0
	}

	queue := make([]*Asix, 0)
	row, col := len(grid), len(grid[0])

	var res int

	bfs := func(b [][]byte, i1, i2 int) {
		queue = append(queue, &Asix{i1, i2})
		for len(queue) > 0 {
			for _, v := range queue {

				queue = queue[1:]

				if inArea(grid, v.x, v.y) && grid[v.x][v.y] == '1' {
					grid[v.x][v.y] = '2'
					queue = append(queue, &Asix{v.x + 1, v.y})
					queue = append(queue, &Asix{v.x - 1, v.y})
					queue = append(queue, &Asix{v.x, v.y + 1})
					queue = append(queue, &Asix{v.x, v.y - 1})
				}

			}
		}

	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if grid[i][j] == '1' {
				bfs(grid, i, j)
				res++
			}

		}
	}
	return res

}

type Asix struct {
	x, y int
}

// numIslands 岛屿数量
func numIslands2(grid [][]byte) int {
	// dfs 扫描二维数组，碰到val = '1' 的情况，
	// 将其标记为2 防止出现重复遍历 (也有标记为0的做陆地「沉没」为海洋)
	// 进行上下左右的dfs

	row, col := len(grid), len(grid[0])

	var dfs func([][]byte, int, int)

	dfs = func(grid [][]byte, r, c int) {

		// 判断base, 不在区域内返回
		if !inArea(grid, r, c) {
			return
		}

		// 如果这个格子不是岛屿，直接返回
		if grid[r][c] != '1' {
			return
		}
		// 将格子标记为「已遍历过」
		grid[r][c] = '2'

		// 周围遍历
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}

	res := 0

	// 遍历网格
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if grid[i][j] == '1' {
				res += 1
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func inArea[T any](grid [][]T, r, c int) bool {

	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])

}
