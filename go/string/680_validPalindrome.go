package string

// validPalindrome 最多可以删除一个字符
func validPalindrome(s string) bool {

	// 双指针(左右指针)
	// 如果左右指针对应的值相同，正常移动
	// 否则，判断左右指针之内的字符是否为回文字符串
	l, r := 0, len(s)-1

	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			flag1, flag2 := true, true
			// 删除r 所在的字符
			for i, j := l, r-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}

			// 删除l 所在的字符
			for i, j := l+1, r; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}

			return flag1 || flag2
		}

	}
	return true
}
