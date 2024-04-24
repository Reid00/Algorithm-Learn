package dptwodim

// isInterleave 交错字符串
func isInterleave(s1 string, s2 string, s3 string) bool {
	// dp[i][j]	表示s1 的前i个和s2 的前j个字符，能否构成s3的i+j 个字符
	// 前i个字符，对应的下标是i-1, 前j个字符是j-1, 第 i+j 个字符是 i + j -1
	// dp[i][j] = (dp[i-1][j] && s1[i-1]==s3[i+j-1]) or (dp[i][j-1] && s2[j-1] = s3[i+j-1])

	m, n := len(s1), len(s2)

	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 初始化
	// 前0个肯定是可以交错拼接而成的
	dp[0][0] = true

	// 第一列初始化
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	// 第一行初始化
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[m][n]
}
