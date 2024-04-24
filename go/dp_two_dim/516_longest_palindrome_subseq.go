package dptwodim

// longestPalindromeSubseq 最长回文子序列 的长度 (文子串是要连续的，回文子序列可不是连续的！)
func longestPalindromeSubseq(s string) int {
	// dp[i][j] 表示s[i,j] 最长回文子序列的最大长度是dp[i][j] (包含i,j)
	// if s[i]==s[j], dp[i][j] = dp[i+1][j-1] + 2
	// if s[i] != s[j], 说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子序列的长度，那么分别加入s[i]、s[j]看看哪一个可以组成最长的回文子序列。
	// 1. 加入s[i], dp[i][j] = dp[i][j-1]
	// 2. 加入s[j], dp[i][j] = dp[+1][j]
	// so. dp[i][j] = max(dp[i][j-1], dp[i+1][j])
	// thus, dp[i][j] = dp[i+1][j-1] +2 Or max(dp[i][j-1], dp[i+1][j])

	n := len(s)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// dp 初始化
	// 1. 因为dp[i][j] 的公式中，不能出现i==j 的情况，这个要单独初始化为dp[i][i] =1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// 因为(i,j) 依赖 (i+1,j-1), (i+1,j), (i, j-1) 分别在左下，正下，左边
	// 因此遍历时要从下到上，才能保证 这三个依赖项先与(i,j) 初始化
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}
