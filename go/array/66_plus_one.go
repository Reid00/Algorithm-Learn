package array

func plusOne(digits []int) []int {

	var carry int

	for i := len(digits) - 1; i >= 0; i-- {

		sum := digits[i] + 1
		digits[i] = sum % 10 // 求余
		carry = sum / 10     // 进位

		if carry <= 0 {
			break
		}
		if i == 0 && carry > 0 {
			// 在第0位插入
			digits = append([]int{carry}, digits...)
		}

	}

	return digits
}
