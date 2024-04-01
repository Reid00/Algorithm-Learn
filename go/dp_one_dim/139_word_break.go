package dponedim

import "slices"

// wordBreak 单词拆分
func wordBreak(s string, wordDict []string) bool {

	// dp[i] 表示s[0:i] 是否在wordDict 中
	// s: leetcode, wordDict: [leet, code]
	// dp[4]=true => s[:4] leet 在wordDict 中
	// dp[i]=true and dp[j] =true => s[i:j] 在wordDict中

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if dp[i] && slices.Contains(wordDict, s[i:j]) {
				dp[j] = true
			}
		}
	}

	return dp[n]
}
