package dptwodim

// longestPalindrome 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串
// Time: O(n2), Space: O(n2)
func longestPalindrome(s string) string {
	// dp[i][j] 表示s[i:j] 范围内是否是一个回文串
	// 从中间向两边扩散，dp[i][j] 依赖(i+1,j-1) 是否是 即:
	// dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	// 因为i,j 依赖( i + 1, j-1) 即是二维数组的左下角，遍历顺序从下向上，做左到右

	n := len(s)
	if n < 2 {
		return s
	}
	// 最长回文子串的起始位置和长度
	maxStart, maxLen := 0, 0

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 遍历顺序很重要，因为(i,j) 依赖(i+1,j-1) 左下角，遍历顺序从下向上，做左到右
	for i := n - 1; i >= 0; i-- {
		// 由于dp 的定义，j >=i 才有意义，所以只需要填充有上边的部分
		for j := i; j < n; j++ {

			if s[i] == s[j] {
				if j-i < 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxStart = i
				maxLen = j - i + 1
			}
		}
	}

	return s[maxStart : maxStart+maxLen]
}

// 中心扩散算法
// Time: O(n2), Space: O(1)
func longestPalindrome2(s string) string {

	if s == "" {
		return s
	}

	l, r := 0, 0

	for i := 0; i < len(s); i++ {

		s1, e1 := expand(s, i, i)   // 以i 为中心向两边扩散
		s2, e2 := expand(s, i, i+1) // 以i,i+1 为中心向两边扩散

		if e1-s1 > r-l {
			l, r = s1, e1
		}

		if e2-s2 > r-l {
			l, r = s2, e2
		}
	}

	return s[l : r+1]
}

// 中心向两边扩散
func expand(s string, i, j int) (start, end int) {

	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	start = i + 1
	end = j - 1
	return
}
