package number

import "math"

// 整数反转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

// 假设环境不允许存储 64 位整数（有符号或无符号）。

func reverse(x int) int {

	var reversed int

	for x != 0 {
		if reversed > math.MaxInt32 || reversed < math.MinInt32 {
			return 0
		}

		// reverse
		// 最后个位数字
		digit := x % 10
		// 进位
		x = x / 10
		reversed = reversed*10 + digit

	}

	return reversed
}
