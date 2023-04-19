package string

import "math"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口思想

	if len(s) == 0 {
		return 0
	}

	// 窗口存放已有的元素和个数
	win := make(map[byte]int)

	// 窗口的左右边界
	var l, r int

	// 长度
	length := math.MinInt64

	for r < len(s) {

		ch := s[r]
		r++
		win[ch]++

		for win[ch] > 1 { // 出现重复的字符, 循环缩短边界。窗口左边界 向右移动
			c := s[l]
			l++
			win[c]--
		}

		if length < r-l {
			length = r - l
		}
	}
	return length
}

func lengthOfLongestSubstring2(s string) int {
	// 窗口里面存放不重复的字串
	// 如果碰到重复的字串，则计算当前的字串长度，并且重置窗口的left

	win := make(map[rune]int)

	var l, resLen int

	for r, ch := range s {

		if _, ok := win[ch]; ok && r >= l { // win 中存在字符ch，并且窗口r 边界大于等于l
			// 记录长度
			resLen = max(resLen, r-l)
			// 重置l 边界, win[ch] 第一次出现ch字符的坐标
			l = win[ch] + 1
		}

		// 将当前字符的idx 记录到win中
		win[ch] = r
	}

	resLen = max(resLen, len(s)-l)
	return resLen
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
