package dptwodim

// minPathSum 最小路径和
// 每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {

	// dp[i][j] 表示走到格子i,j 位置的，最小路径和
	// 到i,j 这个位置，可以从(i,j-1) 或者(i-1,j) 移动而来
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]

	n := len(grid)
	m := len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	// 初始化
	dp[0][0] = grid[0][0]
	// 最左侧的列
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 最上面的行
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]
}
