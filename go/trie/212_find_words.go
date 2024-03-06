package trie

// findWords 单词搜索 II
func findWords(board [][]byte, words []string) []string {

	// 1. 字典树 insert word中的单词
	// 2. 遍历board 的每个字符时，dfs 上下左右的孩子，组成单词查看是否在words 中
	// 3. 同一个单元格内的字母在一个单词中不允许被重复使用, 记录下，单个dfs 中防止被重复使用

	// 生成字典树
	trie := Constructor()
	for _, w := range words {
		trie.Insert(w)
	}

	var ret = make([]string, 0)

	// 单个字符一个单词中只能使用一次
	// visited 保证一个搜索的DFS 中，该字符只被使用一次
	m, n := len(board), len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 邻居
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	// 遍历
	var dfs func(board [][]byte, r, c int, t *Trie)
	dfs = func(board [][]byte, r, c int, t *Trie) {
		// 无效坐标
		if !inArea(board, r, c) {
			return
		}

		// 提前剪枝
		if t == nil || t.count == 0 {
			return
		}

		// 字符已经被使用过
		if visited[r][c] {
			return
		}

		visited[r][c] = true

		// 此时字符是单词终点
		if t.end {
			ret = append(ret, t.word)
			// 整个Trie去剪枝
			trie.Del(t.word)
		}

		// 向周围邻居节点DFS
		for _, v := range dirs {
			x, y := r+v[0], c+v[1]
			if inArea(board, x, y) {
				idx := board[x][y] - 'a'
				dfs(board, x, y, t.children[idx])
			}
		}

		// 恢复现场
		visited[r][c] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := board[i][j] - 'a'
			dfs(board, i, j, trie.children[idx])
		}
	}
	return ret
}

func inArea(board [][]byte, r, c int) bool {
	return r >= 0 && r < len(board) && c >= 0 && c < len(board[0])
}
