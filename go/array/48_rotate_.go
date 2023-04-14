package array

// rotate 旋转图像
func rotate2(matrix [][]int) {

	// 水平翻转 + 对角线翻转  == 顺时针旋转90°

	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
