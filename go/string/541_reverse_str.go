package string

// reverseString2
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，
// 就反转这 2k 字符中的前 k 个字符。

// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
func reverseStr(s string, k int) string {

	n := len(s)
	s1 := []byte(s)

	// 直接模拟
	for i := 0; i < n; i += 2 * k {
		sub := s1[i:min(i+k, n)]
		for j, h := 0, len(sub); j < h/2; j++ {
			sub[j], sub[h-1-j] = sub[h-1-j], sub[j]
		}
	}
	return string(s1)

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
