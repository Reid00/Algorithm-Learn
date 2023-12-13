package array

func lengthOfLastWord(s string) int {
	// 从后往前遍历，遇到单词之后第一个空白字符停止
	idx := len(s) - 1
	ans := 0

	for s[idx] == ' ' {
		idx--
	}

	for idx >= 0 && s[idx] != ' ' {
		ans++
		idx--
	}
	return ans
}
