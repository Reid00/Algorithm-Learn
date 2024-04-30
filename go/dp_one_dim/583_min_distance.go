package dponedim

// minDistance2 两个字符串的删除操作  不同于 72. 编辑距离 此处只能删除
// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
func minDistance2(word1 string, word2 string) int {
	// dp[i][j] 表示word1 以i-1 为下标的字符串，转化为以j-1为下标的字符串word2,最小删除的步数是dp[i][j]
	// if word1[i-1] == word2[j-1] => dp[i][j] = dp[i-1][j-1]
	// if word1[i-1] != word2[j-1] => dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1)
	// word1 前i-1个下标中删除一个字符，转化为word2[:j] 或者word1不懂，word2前j-1 个字符删除一个字符转为word1[:i]

	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i < m+1; i++ {
		// word1[:i] 前i个字符转化为空字符串，需要i步，删除i 的全部字符
		dp[i][0] = i
	}

	for j := 0; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}

	return dp[m][n]
}
