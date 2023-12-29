package string

// findSubstring 串联所有单词的子串
func findSubstring(s string, words []string) []int {

	// 滑动窗口思想
	// words 中的单词长度是k, 则循环k次 看s窗口中的长度k 是否在words 中存在

	if len(words) == 0 {
		return []int{}
	}

	k := len(words[0])
	n := len(words)
	sLen := len(s)

	// s 的长度 不如 words 拼接之后的长度
	if sLen < k*n {
		return []int{}
	}

	var ans = make([]int, 0)

	// k:v => word: cnt 用次来判断窗口是否满足
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}

	// 循环k 次，每次步长k 相当于把所有的s 中的组合都尝试了
	for i := 0; i < k; i++ {
		// 记录窗口中已经满足单词的个数，if count == n  则记录
		count := 0
		// 记录窗口中每个单词的个数
		wordCnt := make(map[string]int)

		for l, r := i, i; r <= sLen-k; r = r + k {
			// r == sLen-k 才有可能窗口匹配到s 中的最后一个字符
			word := s[r : r+k]

			// 单词在 wordMap 中存在
			if num, ok := wordMap[word]; ok {
				// 窗口中的cnt 要超过了(数目相等)，右移 窗口左指针
				for wordCnt[word] >= num {
					wordCnt[s[l:l+k]]--
					count--
					l += k
				}

				wordCnt[word]++
				count++

			} else {
				// 单词在 wordMap 中不存在时，该窗口无效，l 到r 的下一个位置
				for l < r {
					wordCnt[s[l:l+k]]--
					count--
					l += k
				}
				// 让l 到r 的下一个位置
				l += k
			}

			if count == n {
				ans = append(ans, l)
			}
		}

	}

	return ans

}
