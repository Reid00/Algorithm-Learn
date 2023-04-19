package array

// setZeroes 矩阵置零
// 空间复杂度 O(m+n)
func setZeroes(matrix [][]int) {

	type pos struct {
		x, y int
	}

	xy := make([]*pos, 0)

	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				// 需要注意的是，你修改后的0 不能作为置零的参考
				// 第一步 先找出 0 的 i，j 然后再 根据i,j 置零
				coordinate := &pos{
					i, j,
				}
				xy = append(xy, coordinate)

				// FIX 下面修改后的0 将会影响置零的条件
				// // 修改行i数据
				// for m := 0; m < len(matrix[0]); m++ {
				// 	matrix[i][m] = 0
				// }
				// // 修改列j 数据
				// for n := 0; n < len(matrix); n++ {
				// 	matrix[n][j] = 0
				// }
			}
		}
	}

	for _, s := range xy {
		for i := 0; i < len(matrix); i++ {
			matrix[i][s.y] = 0
		}

		for j := 0; j < len(matrix[0]); j++ {
			matrix[s.x][j] = 0
		}
	}

}

// setZeroes2 O(1) 的额外空间
// 如果某行，某列出现0， 只需要将第一行，第一列的数据标记为零
// 关键思想: 用matrix第一行和第一列记录该行该列是否有0,作为标志位
func setZeroes2(matrix [][]int) {

	n, m := len(matrix), len(matrix[0])
	// 用两个额外元素来标记第一行和第一列是否有0
	// 否则其他行列有0 的数据都归到第一行，第一列，导致无法判断第一行，第一列是否有0 了
	row0, col0 := false, false

	// 先判断第一行，第一列
	// 第一列
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	for i := 0; i < m; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	// 遍历其他行列，如果存在0， 将0 置于首行，首列之中
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 重新遍历首行元素讲对应的行列 置于0
	// 开始遍历第一行，第一列的时候 要跳过左上角的第一个元素，
	// 因为他会把首行，首列置为0， 后续遍历就无效了
	// 第一列
	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			// 将i 行的全部列元素置为 0
			// j 从0或者1 开始都行
			for j := 1; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}
	// 第一行
	for i := 1; i < m; i++ {
		if matrix[0][i] == 0 {
			// 将i 列的全部行元素置为 0
			for j := 1; j < n; j++ {
				matrix[j][i] = 0
			}
		}
	}

	// 判断首行，首列
	if col0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
	if row0 {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}

}
