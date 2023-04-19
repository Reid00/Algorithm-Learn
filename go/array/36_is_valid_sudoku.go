package array

// isValidSudoku 9 x 9 的数独是否有效
func isValidSudoku(board [][]byte) bool {

	// 记录元素e 出现在数组的哪一行，哪一列
	var rows, cols [9][9]int

	var subboxs [3][3][9]int

	for i, row := range board {
		for j, c := range row {
			// byte 类型 二维数组，里面是字符
			if c == '.' {
				continue
			}
			// 字符的内容是1-9， 用长度9 的数组存储，index 减一
			index := c - '1'

			// 例如如果为坐标(1,1)为8， 则在rows[i][8-1]的位置 +1
			rows[i][index]++
			cols[j][index]++
			subboxs[i/3][j/3][index]++

			if rows[i][index] > 1 || cols[j][index] > 1 || subboxs[i/3][j/3][index] > 1 {
				return false
			}

		}
	}

	return true
}
