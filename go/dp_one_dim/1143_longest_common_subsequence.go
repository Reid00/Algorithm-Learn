package dponedim

// longestCommonSubsequence 最长公共子序列  和 718 最长重复数组较像，只不过此处可以不必连续
func longestCommonSubsequence(text1 string, text2 string) int {
	// 同 718，dp[i][j] 表示text1的下标i-1 和 text3 的下标j-1，之间最长公共子序列的长度是dp[i][j]
	// 状态转移方程 dp[i][j] 的值有两种情况
	// 1. text1[i-1] == text1[j-1], 那么dp[i][j] = dp[i-1][j-1] + 1
	// 2. text1[i-1] != text1[j-1], 则看text1[0..i-2] 和 text2[0..j-1] 的公共子序列
	// 或者text1[0..i-1] 和text2[0..j-2] 的值，去两者的最大值
	// dp[i][j] = max(dp[i-1][j], dp[i][j-1])

	m, n := len(text1), len(text2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始化；因为dp[i][j] 是表示下标(i-1,j-1) 的 最长公共子序列，所以i/j == 0 都是无意义的
	// 可以初始化为0

	ans := 0

	// 状态转移方程
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

			ans = max(ans, dp[i][j])
		}

	}

	return ans
}
