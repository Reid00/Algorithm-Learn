package array

import (
	"fmt"
	"math"
)

/* *
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
* */

func reverse(x int) int {
	var rev int

	for x != 0 {
		// rev 压栈（*10）前先判断是否越界
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}

		// 出栈
		// 123 为例, 获取最低位数字
		digit := x % 10
		// 去掉最低位的后x的数字
		x /= 10

		// 压栈: rev * 10 + 最低位
		rev = rev*10 + digit
		fmt.Printf("digit: %d, input x: %d, rev: %d\n", digit, x, rev)
	}
	return rev
}
