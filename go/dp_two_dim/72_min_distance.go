package dptwodim

// minDistance 编辑距离
func minDistance(word1 string, word2 string) int {
	// dp[i][j] word1的前i个字符和word2的前j个字符的编辑距离。意思就是word1的前i个字符，变成word2的前j个字符，最少需要这么多步。
	// if word1[i-1] == word2[j-1]: dp[j][j] = dp[i-1][j-1]
	// if word1[i-] != word2[j-1]:
	// 有三种操作，1. 删: word1删除最后一个, dp[i][j] = dp[i-1][j] + 1
	// => word1删除一个元素，那么就是以下标i - 2为结尾的word1 与 j-1为结尾的word2的最近编辑距离 再加上一个操作
	// 删: word2 中删一个, dp[i][j] = dp[i][j-1] + 1
	// 2. 增： 增的本质和删一样，word1 删除一个，相当于word2 增加一个
	// 3. 改: word1替换word1[i - 1]，使其与word2[j - 1]相同，此时不用增删加元素, 只需要一次替换即可在i-2为下标的基础上转化
	// 即: dp[i][j] = dp[i-1][j-1] + 1
	// 总结, if word1[i-1]!=word2[j-1]=> dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	dp[0][0] = 0
	// 初始化第一列
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	// 初始化第一行
	for j := 1; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}
