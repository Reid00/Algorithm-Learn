package graph

// largestIsland 最大人工岛
func largestIsland(grid [][]int) int {
	// DFS
	// 第一遍，遍历出所有岛屿，并且给岛屿编号，同时面积记录到hmap 中 {id: area}
	// 第二遍，遍历出海洋(val 0), 找到周围的邻居岛屿，切记不能重复
	// if 有, 面积=1 + 相邻岛屿的面积和

	// 最大面积
	var res int
	// dfs 求面积, 并且给岛屿标注idx
	var dfs func([][]int, int, int, int) int

	dfs = func(grid [][]int, r, c int, flag int) int {

		// 不在格子内
		if !inArea(grid, r, c) {
			return 0
		}

		// 不是岛屿
		if grid[r][c] != 1 {
			return 0
		}

		// 标记已经遍历的个格子
		grid[r][c] = flag

		return 1 + dfs(grid, r-1, c, flag) +
			dfs(grid, r+1, c, flag) +
			dfs(grid, r, c+1, flag) +
			dfs(grid, r, c-1, flag)
	}

	hmap := make(map[int]int)
	// 岛屿的索引从2 开始
	var idx = 2

	row, col := len(grid), len(grid[0])

	// 第一次遍历，求岛屿面积并且赋给岛屿索引，并记录面积
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, i, j, idx)
				res = max(res, area)
				hmap[idx] = area
				idx++
			}
		}
	}

	// 如果一块岛屿都没有返回1, 翻转一个
	if res == 0 {
		return 1
	}

	// 再次遍历，找到海洋，判断周围是否有岛屿
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 0 {

				hset := findNeighbor(grid, i, j)

				// 周围没有岛屿
				if len(hset) == 0 {
					continue
				}

				// idx 是岛屿的索引
				total := 0
				for idx, _ := range hset {
					// 周围岛屿的面积和
					total += hmap[idx]
				}
				res = max(res, total+1)
			}
		}
	}

	return res
}

// findNeighbor 找到周围有哪几个岛屿{idx: true}
func findNeighbor(grid [][]int, x, y int) map[int]bool {

	// 之所以用set 不用array 防止加重复, 周围可能是同一个岛屿id
	hset := make(map[int]bool)

	// 索引有效，并且上面是岛屿
	if inArea(grid, x-1, y) && grid[x-1][y] != 0 {
		hset[grid[x-1][y]] = true
	}
	// 索引有效，并且下面是岛屿
	if inArea(grid, x+1, y) && grid[x+1][y] != 0 {
		hset[grid[x+1][y]] = true
	}
	// 索引有效，并且左侧是岛屿
	if inArea(grid, x, y-1) && grid[x][y-1] != 0 {
		hset[grid[x][y-1]] = true
	}
	// 索引有效，并且右侧是岛屿
	if inArea(grid, x, y+1) && grid[x][y+1] != 0 {
		hset[grid[x][y+1]] = true
	}

	return hset
}
