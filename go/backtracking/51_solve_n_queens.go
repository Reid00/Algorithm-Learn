package backtracking

import (
	"strings"
)

// solveNQueens N皇后
func solveNQueens(n int) [][]string {

	// 搜索中满足条件的结果
	path := make([][]string, n)
	for i := 0; i < n; i++ {
		path[i] = make([]string, n)
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}

	// 所有满足条件的结果集
	res := make([][]string, 0)

	// 回溯算法
	var backtrack func(row int)
	backtrack = func(row int) {
		// 遍历到最后一行了
		if row == n {
			tmp := make([]string, 0)
			for i := 0; i < n; i++ {
				tmp[i] = strings.Join(path[i], "")
			}
			res = append(res, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col, n, path) {
				path[row][col] = "Q"
				backtrack(row + 1)
				// 回溯
				path[row][col] = "."
			}
		}
	}

	backtrack(0)
	return res
}

// isValid path[r][c] 是否可以设置为Q
// 同行，同列，对角线的都为无效
func isValid(r, c int, n int, path [][]string) bool {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if path[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}
