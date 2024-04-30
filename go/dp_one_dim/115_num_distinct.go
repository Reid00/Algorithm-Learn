package dponedim

// numDistinct 不同的子序列		类似392题
// s 序列中 t 出现的个数
func numDistinct(s string, t string) int {
	// dp[i][j] 表示以i-1 下标为结尾的s序列中，出现以j-1 下标为结尾的子序列t 的个数是dp[i][j]
	// if s[i-1] == t[j-1]，有两种可能，不管最后一个字符 dp[i][j] = dp[i-1][j-1]
	// 第二个情况是，在s 中不用以i-1 下标为结尾，删除一个元素继续匹配，可以用i-2 即dp[i][j] = dp[i-1][j]
	// s 中前i-2 个字符中，有多少个t 中前j-1 个字符, 比如s=bagg t=bag dp[4][3] 可以用s[0][1][2] 匹配t[3]
	// so dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// if s[i-1] != t[j-1], dp[i][j] = dp[i-1][j]

	m, n := len(s), len(t)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i < m+1; i++ {
		// s 以i-1 结尾的子串，只有删除所有i-1个字符才能转变为空字符串，方法只有1
		dp[i][0] = 1
	}

	for j := 0; j < n+1; j++ {
		// s 空字符串删除几个字符才能转变为t[:j-1], 方法只有0
		dp[0][j] = 0
	}
	// 空字符串s，可以删除0个元素，变成空字符串t
	dp[0][0] = 1

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
