package string

import "math"

// reverse 反转int32
func reverse(x int) int {

	// 实际上我们只要能拿到这个整数的 末尾数字 就可以了。
	var ret int

	for x != 0 {
		if x < math.MinInt32 / 10 || x > math.MaxInt32 / 10 {
			return 0
		}

		carry := x % 10
		x /= 10
		ret = ret * 10 + carry
	}
	return ret
}