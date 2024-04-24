package dptwodim

// uniquePathsWithObstacles 不同路径 II
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
// 网格中的障碍物和空位置分别用 1 和 0 来表示。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	// dp[i][j]	表示到达(i,j) 位置总共有dp[i][j] 条路径
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]

	n, m := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 初始化dp
	dp[0][0] = 0

	// 最左侧的列
	for i := 0; i < n; i++ {
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
			// 后续也都为0
			for x := i + 1; x < n; x++ {
				dp[x][0] = 0
			}
			break
		}

	}
	// 最上面的行
	for j := 0; j < m; j++ {
		if obstacleGrid[0][j] != 1 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
			for x := j + 1; x < m; x++ {
				dp[0][x] = 0
			}
			break
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	return dp[n-1][m-1]
}
