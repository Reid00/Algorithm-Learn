package math

// plusOne 加一
func plusOne(digits []int) []int {

	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		// 加1
		digits[i] += 1
		// 求余 为对应i位置
		digits[i] = digits[i] % 10

		// 末位+1 后 如果不为10，则直接返回，否则前一位+1
		if digits[i] != 0 {
			return digits
		}
	}

	// digits 中所有的元素均为 9
	digits = make([]int, n+1)
	digits[0] = 1

	return digits
}
