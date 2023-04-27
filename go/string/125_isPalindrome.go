package string

import (
	"strings"
	"unicode"
)

// isPalindrome 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
// 短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
// O(n) 空间复杂度
func isPalindrome(s string) bool {

	var str strings.Builder

	str.Grow(len(s))

	// 大写转小写
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			ch |= 32
		}
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			str.WriteRune(ch)
		}
	}

	l := str.String()

	for i := 0; i < len(l); i++ {
		if l[i] != l[len(l)-1-i] {
			return false
		}
	}
	return true

}

// isPalindrome2 O(1) 空间复杂度  双指针(左右指针)
func isPalindrome2(t string) bool {
	s := strings.ToLower(t)

	l, r := 0, len(t)-1
	for l < r {
		for l < r && !isalnum(s[r]) {
			r--
		}

		for l < r && !isalnum(s[l]) {
			l++
		}

		// left right 移动之后，先判断大小关系
		if l < r {
			if s[l] != s[r] {
				return false
			}

			l++
			r--
		}
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
