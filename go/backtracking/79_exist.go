package backtracking

// exist 单词搜索
func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}

	// visited 表示该位置是否被访问过
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	// 邻居
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	// r,c 表示在board 中的坐标， idx 表示将要匹配的word 中的idx
	var backtrack func(r, c, idx int) bool
	backtrack = func(r, c, idx int) bool {
        if board[r][c] != word[idx] { // 剪枝: 当前字符不匹配
            return false
        }

		// 已经匹配到最后一个单词量了
		if idx == len(word)-1 {
			return board[r][c] == word[idx]
		}

		if board[r][c] == word[idx] {
			visited[r][c] = true

			for _, v := range dirs {
				x, y := r+v[0], c+v[1]
				if inArea(board, x, y) && !visited[x][y] {
					if backtrack(x, y, idx+1) {
						return true
					}
				}
			}
			// 回溯
			visited[r][c] = false
		}
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// inArea 坐标 (i,j) 是否在的区域内
func inArea(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}
