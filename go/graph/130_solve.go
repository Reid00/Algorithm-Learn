package graph

// solve 被围绕的区域
// 只修改被X包围的O
func solve(board [][]byte) {

	// Graph 的惯用套路，dfs
	if len(board) == 0 {
		return
	}

	row, col := len(board), len(board[0])

	// 碰到在边缘的O,进行DFS，把这些O 连接的O 标记为A，这些是不需要被修改的
	var dfs func([][]byte, int, int)

	dfs = func(b [][]byte, i1, i2 int) {

		if !inArea(b, i1, i2) {
			return
		}

		if b[i1][i2] != 'O' {
			return
		}
		// 讲O 标记为A
		b[i1][i2] = 'A'

		dfs(b, i1+1, i2)
		dfs(b, i1-1, i2)
		dfs(b, i1, i2+1)
		dfs(b, i1, i2-1)
	}

	// 遍历边缘的O
	// 遍历最左侧和右侧的列
	for i := 0; i < row; i++ {
		// 加不加这个==O 的判断影响不大
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][col-1] == 'O' {
			dfs(board, i, col-1)
		}
	}
	// 遍历最上面和下面的行
	for i := 0; i < col; i++ {
		dfs(board, 0, i)
		dfs(board, row-1, i)
	}

	// 找到剩余没有 被修改为A 的O，直接修改为X
	// 同时碰到A 改回O
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}

}
