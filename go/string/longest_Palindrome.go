package string

// 给你一个字符串 s，找到 s 中最长的回文子串。
// no5, 最长回文子串

// dp
// dp[i][j] 表示s[i..j] 是否位回文子串
// 如果dp[i][j] 是，那么dp[i+1][j-1]也是
// dp[i][j] = {
//     dp[i+1][j-1]  && dp[i]==dp[j]	case 3
//     dp[i][i]    一个字符				case 2
//     dp[i][i+1]  两个字符相同时		case 1
// }
func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))

	result := s[0:1] // 初始化结果，最小的回文就是单个字符

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // case 2
	}

	for length := 2; length <= len(s); length++ {
		// 固定长度，寻找组合
		for start := 0; start < len(s)-length+1; start++ {
			// 长度固定的情况下，不断的移动start 的值，遍历回文的可能性
			// end - start + 1 = length
			end := length + start - 1
			// 判断是否为回文
			if s[end] != s[start] {
				continue
			} else if length < 3 {
				// case 1
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1] // case 3
			}

			// 根据长度，更新result
			if dp[start][end] && (end-start+1) > len(result) {
				result = s[start : end+1]
			}

		}
	}

	return result
}

// 中心扩散算法
// 从每个字符吵两边扩散，记录为回文的情况
func longestPalindrome2(s string) string {

	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		left1, right1 := expandFromCenter(s, i, i)
		left2, right2 := expandFromCenter(s, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}

		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandFromCenter(s string, left, right int) (int, int) {

	// 一直向两边扩散，知道不满足回文
	for ; s[left] == s[right] && left >= 0 && right <= len(s)-1; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
