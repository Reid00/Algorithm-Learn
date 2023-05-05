package string

import (
	"strings"
)

// myAtoi  result = result * 10 + i
// 需要注意int32 边界问题  + 符号 正负号问题
func myAtoi(s string) int {

	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	var negative bool

	if s[0] == '-' {
		negative = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var result int

	for _, ch := range s {
		// 字符是数字
		if ch >= '0' && ch <= '9' {
			// 距离字符0 的距离，即为数字
			ch -= '0'
			result = result*10 + int(ch)

			if result >= 1<<31 {
				result = 1 << 31
				break
			}
		} else {
			// 非数字 字符则退出循环
			break
		}

	}

	// 处理边界问题
	// 数字范围[-2 ** 31, 2 ** 31 -1]
	if negative {
		result = -result
		// 默认现在最大是2 ** 31 次方，负数直接加符号即可
		// if result <= -(1 << 31) {
		// 	result = -1 << 31
		// }
	} else {
		if result >= 1<<31 {
			result = 1<<31 - 1
		}
	}
	return result
}
