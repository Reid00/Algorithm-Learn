package backtracking

// totalNQueens N 皇后II
// 返回 n 皇后问题 不同的解决方案的数量。
func totalNQueens(n int) int {

	// 初始化棋盘，并作为搜索路径
	path := make([][]string, n)
	for i := 0; i < n; i++ {
		path[i] = make([]string, n)
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}
	// 满足条件的集合数量
	res := 0

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res += 1
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col, n, path) {
				path[row][col] = "Q"
				backtrack(row + 1)
				// backtrack
				path[row][col] = "."
			}
		}
	}

	backtrack(0)
	return res
}
