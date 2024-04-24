package dptwodim

import "math"

// minimumTotal 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	// dp[i][j] 表示第i行，j列的 最小路径和为 dp[i][j], i,j 从0 开始算的
	// dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]

	n := len(triangle)

	// defenition dp
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// init dp
	// 由于dp[i][j] 和 上一层的i-1 有关，j和 j-1或者j有关, dp[0][0] = triangle[0][0]
	dp[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		// 最左列
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		// 最右列
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := math.MaxInt32
	// 最行一行中取到最小值
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}
