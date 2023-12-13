package string

import "fmt"

// strStr 请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标
// （下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
func strStr(haystack string, needle string) int {
	for i, ch := range haystack {
		if byte(ch) == needle[0] {
			// 曾经忘记加下面的判断
			if i+len(needle) > len(haystack) {
				return -1
			}
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}

// strStr2 KMP 算法, 详情: https://www.cnblogs.com/zzuuoo666/p/9028287.html
// https://www.zhihu.com/question/21923021/answer/281346746
// TODO to understand
func strStr2(haystack string, needle string) int {

	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	// needle[0:j] 最大公共前缀后缀表
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}

	fmt.Println("next: ", pi)

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == m {
			return i - m + 1
		}

	}
	return -1
}
