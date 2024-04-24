package dptwodim

// countSubstrings 回文子串  数量
// 中心扩展算法 O(n2) 的时间复杂度 + O(1) 的空间
func countSubstrings(s string) int {

	n := len(s)
	res := 0
	// 中心扩展算法有一个需要注意的是，开始的中心可以是一个，比如aba 中的b, 也可以是两个字符abba 中的bb
	for i := 0; i < n; i++ {
		res += expandFromCenter(s, i, i)   // 以i为中心
		res += expandFromCenter(s, i, i+1) // 以i 和 i+1 为中心
	}

	return res
}

func expandFromCenter(s string, i, j int) int {
	res := 0

	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
		res++
	}
	return res
}

// countSubstrings 回文子串  数量
// dp O(n2) 的时间复杂度 + O(n2) 的空间复杂度
func countSubstrings2(s string) int {
	// dp[i][j] 表示子串s[i,j] 是否是回文子串, 包含i,j (j>=i)
	// dp[i][j] 依赖其更短的子串 s[i+1,j-1]是否是回文子串
	// dp[i][j] = dp[i+1][j-1] && s[i]==s[j]
	// 由于(i,j) 依赖(i+1, j-1) 左下角的元素，所以遍历时，不能常规从左到右，从上到下
	// 应该从下到上，从左到右

	n := len(s)
	res := 0

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		// 由于dp 的定义，j >= i 才会有意义，所以只需要填充上半部分即可
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i < 2 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = dp[i+1][j-1]
					res++
				}
			}
		}
	}

	return res
}
