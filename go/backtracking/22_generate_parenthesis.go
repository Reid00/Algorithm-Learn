package backtracking

import "strings"

// generateParenthesis 括号生成
func generateParenthesis(n int) []string {
	// 回溯搜索

	// 所有满足条件的集合
	res := make([]string, 0)

	// 搜索中满足条件的集合
	path := make([]string, 0)

	// symbol 记录当前左括号的个数, symbol < n
	// path 中push ( symbol + 1; push ） symbol - 1, 最后symbol 被抵消==0
	// idx path 中当前元素下标
	var backtracck func(symbol, idx int)
	backtracck = func(symbol, idx int) {
		if idx == n*2 {
			if symbol == 0 {
				p := strings.Join(path, "")
				res = append(res, p)
			}
			return
		}

		// 可以往path栈里面push (, 每次push 左括号 symbol + 1
		if symbol < n {
			path = append(path, "(")
			backtracck(symbol+1, idx+1)
			path = path[:len(path)-1]
		}
		// 往path 里面push ), 抵消左括号数量
		if symbol > 0 {
			path = append(path, ")")
			backtracck(symbol-1, idx+1)
			path = path[:len(path)-1]
		}
	}

	backtracck(0, 0)
	return res
}
