package backtracking

import (
	"strings"
	"unicode"
)

// letterCasePermutation 字母大小写全排列
func letterCasePermutation(s string) []string {
	// 回溯

	if len(s) == 0 {
		return nil
	}

	// 所有结果的合集
	var res = make([]string, 0)
	// 当前查询的路径合集
	path := make([]string, 0)

	// idx 当前遍历到s 的索引
	var backtrack func(idx int)
	backtrack = func(idx int) {

		if idx == len(s) {
			p := strings.Join(path, "")
			res = append(res, p)
			return

		}

		path = append(path, string(s[idx]))
		backtrack(idx + 1)
		path = path[:len(path)-1]

		if s[idx] >= 'a' && s[idx] <= 'z' {
			path = append(path, string(unicode.ToUpper(rune(s[idx]))))
			backtrack(idx + 1)
			path = path[:len(path)-1]
		} else if s[idx] >= 'A' && s[idx] <= 'Z' {
			path = append(path, string(unicode.ToLower(rune(s[idx]))))
			backtrack(idx + 1)
			path = path[:len(path)-1]
		}

	}

	backtrack(0)
	return res
}
