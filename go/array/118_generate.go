package array

// 杨辉三角
func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		// 每行第一个数字
		result[i][0] = 1
		// 每行最后一个数字
		result[i][i] = 1

		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]

		}
	}

	return result
}
