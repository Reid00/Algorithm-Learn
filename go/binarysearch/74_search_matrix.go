package binarysearch

// searchMatrix 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {

	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1

	for l <= r {
		mid := l + (r-l)>>1

		// 对列求差 和 求余 分别得到行列坐标
		row, col := mid/n, mid%n

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}

	return false

}
