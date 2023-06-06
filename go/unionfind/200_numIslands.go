package unionfind

// numIslands  并查集计算连通分量
func numIslands(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	// 二维数组转化成一维 坐标(x,y) => x*n + y n 是列数
	// dummy 虚拟父节点
	dummy := m * n

	uf := NewUnionFind(dummy + 1)

	// d 常用来表示上下左右四个方向搜索
	d := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 遍历grid 找到 1(陆地) 和 上下左右连通
			if grid[i][j] == '1' {
				for k := 0; k < 4; k++ {
					// 避免x i 是0， 向左搜索减 1， j 是0 向上搜索减 1 的情况
					if i == 0 && d[k][0] == -1 {
						continue
					} else if j == 0 && d[k][1] == -1 {
						continue
					} else if i == m-1 && d[k][0] == 1 {
						// 避免x 是最下面， 向下搜索， j 是最右侧 向右搜索的情况
						continue
					} else if j == n-1 && d[k][1] == 1 {
						continue
					}
					x := i + d[k][0]
					y := j + d[k][1]
					if grid[x][y] == '1' {
						uf.Union(i*n+j, x*n+y)
					}
				}
			} else {
				// 所有的 0 和dummy 相连，记作一个连通分量
				uf.Union(dummy, i*n+j)
			}
		}
	}

	return uf.count - 1
}
