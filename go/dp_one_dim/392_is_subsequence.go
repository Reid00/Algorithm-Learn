package dponedim

// isSubsequence 判断子序列, s 是 t 的子序列
// dp 为72.编辑距离做基础
func isSubsequence(s string, t string) bool {
	// m <=n s 是否 是 t 的子序列
	// dp[i][j] 表示s以i-1 结尾和t 以j-1 结尾的相同最长子序列长度是dp[i][j]
	// if s[i-1] == t[j-1] => dp[i][j] = dp[i-1][j-1] + 1
	// if s[i-1] != t[j-1], 因为要判断s是否是t 的子序列，所以s 的下标不能减少
	// t 的可以缩短， 所以不相同时，t 可以剪短，删除一个字符，dp[i][j] = dp[i][j-1]

	m, n := len(s), len(t)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化, dp[0][0] 无意义，初始化为0
	// dp[i][0] 表示以下标i-1为结尾的字符串，与空字符串的相同子序列长度，所以为0. dp[0][j]同理。
	// 默认是0 无需初始化

	// 状态转移方程 遍历顺序
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// t 字符串删除一个字符
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n] == m
}

// isSubsequence 判断子序列, s 是 t 的子序列
// 双指针
func isSubsequence_(s string, t string) bool {
	// m <=n
	m, n := len(s), len(t)

	idx1, idx2 := 0, 0

	for idx1 < m && idx2 < n {
		if s[idx1] == s[idx2] {
			idx1++
		}
		// t 字符串的指针一定要迁移的
		idx2++
	}

	return m == idx1
}
