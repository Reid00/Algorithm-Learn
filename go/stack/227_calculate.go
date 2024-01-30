package stack

import "unicode"

// calculate 基本计算器 II
func calculate2(s string) int {

	// 因为有 * / 无法知道计算顺序，先把数字压栈，最后 + 得以结果
	// * / 碰到之后先计算结果 再压栈
	// - 压栈 相反数

	prevSign := '+'
	stack := make([]int, 0)

	// 要压栈的数字
	num := 0
	for i, ch := range s {
		// 可能存在多位的整数 "123 + 4"
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
		}

		// 是 运算符
		if !unicode.IsDigit(ch) && ch != ' ' || i == len(s)-1 {

			// 根据 前一位的运算符做计算
			switch prevSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				// 前一位是 * 则此时 乘法的left right 都可以知道
				// left 为 stack[len(stack)-1]
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}

			// 更新前一位运算符
			prevSign = ch
			num = 0
		}
	}

	var res int
	for _, v := range stack {
		res += v
	}

	return res
}
