package string

import "math"

//  minWindow 最小覆盖子串
func minWindow(s string, t string) string {
	// 滑动窗口思想
	// need 是需要的字符，win 是窗口中的字符，都用hashMap 记录
	// 当win中的元素满足条件是，win 左边界l 右移 ++ ,不断观察win 中元素是否满足条件

	win, need := make(map[byte]int), make(map[byte]int)

	for _, ch := range t {
		// range 遍历时 ch 是rune 类型，转成byte 类型
		need[byte(ch)]++
	}

	// 窗口长度的值
	min := math.MaxInt32

	// 窗口的左右边界下标
	l, r := 0, 0

	// win 中的元素是否完全匹配到 need 中的个数了
	match := 0
	// 满足条件下的最小长度时的索引
	start, end := 0, 0

	for r < len(s) {

		ch := s[r]
		r++

		// 这个元素是需要的
		if need[ch] > 0 {
			win[ch]++
			if win[ch] == need[ch] {
				match++
			}
		}

		// 窗口中的元素已经完全匹配到，满足条件了
		// match == len(t) 不能用这个，因为t 中可能有重复元素，那match 值 小于 len(t)
		for match == len(need) {
			// 窗口左边界的字符
			chL := s[l]

			// 先记录下满足条件的窗口的左右边界索引
			if r-l < min {
				start, end = l, r
				min = r - l
			}

			// 如果chL 是需要的，则win中的元素个数需要更新
			if need[chL] > 0 {
				if win[chL] == need[chL] {
					match--
				}
				win[chL]--
			}
			l++
		}

	}

	if min == math.MaxInt32 {
		return ""
	}

	return s[start:end]
}
