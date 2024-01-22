package array

// spiralOrder 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	// 模拟
	// 行，列 (m,n)
	rows, cols := len(matrix), len(matrix[0])

	visited := make([][]bool, rows)
	// 初始化visited 矩阵
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var (
		// 矩阵的总长度
		total = rows * cols
		// 顺时针访问的顺序
		order    = make([]int, total)
		row, col = 0, 0
		// 索引访问的方向，用这个二维数组指示
		directions = [][]int{
			// 列+1， 向右
			{0, 1},
			// 行+1， 向下
			{1, 0},
			// 列-1， 下左
			{0, -1},
			// 行-1， 向上
			{-1, 0},
		}
		directionIdx = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][col]
		visited[row][col] = true

		nextRow, nextCol := row+directions[directionIdx][0], col+directions[directionIdx][1]
		// 判断下一次索引位置是否合法
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			// mod 4 是因为 directionIdx 一定在0-3 之间
			directionIdx = (directionIdx + 1) % 4
		}

		row += directions[directionIdx][0]
		col += directions[directionIdx][1]
	}
	return order
}
