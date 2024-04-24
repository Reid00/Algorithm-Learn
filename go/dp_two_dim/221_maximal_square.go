package dptwodim

// maximalSquare 最大正方形
// dp
// Time: O(m*n), Space: O(m*n)
func maximalSquare(matrix [][]byte) int {
	// dp[i][j] 表示以(i,j)坐标且值包含 1 的正方形的最大边长
	// 状态转移方程:
	// 以(i,j) 为右下角正方形边长受限于其正上方，左上方，左边正方形的边长的最小值
	// dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

	m, n := len(matrix), len(matrix[0])

	// dp 中最大的边长
	maxSide := 0

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		// 初始化
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	// 赋值状态转移方程
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 因为已经根据matrix[i][j] 赋值过了，所以用dp[i][j] 判断即可
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
