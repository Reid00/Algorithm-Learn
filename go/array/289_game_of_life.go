package array

// gameOfLife 生命游戏
func gameOfLife(board [][]int) {
	// 引用复合状态，0 dead, 1 live 如果直接修改0->1 或者 1-> 0 会让修改后的结果影响它后续的判断
	// 所以引入 -1 代表原来是1，会死亡，后续需要变成0 (此处比较经典的是：-1 abs 之后变成1 后续不影响统计下个元素)
	// 引入 2 代表原来是0，复活，后续会变成1

	// 生命游戏的八个方向
	directions := [][]int{
		// 右
		{0, 1},
		// 左
		{0, -1},
		// 上
		{-1, 0},
		// 下
		{1, 0},
		// 左上
		{-1, -1},
		// 右上
		{-1, 1},
		// 右下
		{1, 1},
		// 坐下
		{1, -1},
	}
	// 宽和高
	m, n := len(board), len(board[0])

	for i, rows := range board {
		for j, v := range rows {
			// 坐标i,j  value v

			// 周围1的数量
			nums := 0

			for _, v := range directions {
				nextRow, nextCol := i+v[0], j+v[1]
				if nextRow < m && nextRow >= 0 && nextCol < n && nextCol >= 0 && Abs(board[nextRow][nextCol]) == 1 {
					nums++
				}
			}

			switch {
			// 规则1,3
			case v == 1 && (nums < 2 || nums > 3):
				board[i][j] = -1
			// 规则4
			case v == 0 && nums == 3:
				board[i][j] = 2
			// 规则2, 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
			default:
			}
		}
	}

	// 再次遍历，将-1,2 置为0,1
	for i, rows := range board {
		for j, v := range rows {
			if v > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
