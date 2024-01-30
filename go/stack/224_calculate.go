package stack

import "unicode"

// calculate 基本计算器
func calculate(s string) int {
	// 栈
	// 碰到( 先把之前的结果保存，优先计算括号内的结果

	var res int
	sign := 1 // +/- 号，初始是1

	n := len(s)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {

		ch := s[i]
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// 如果是多位数先表示其数学含义 如 "123+1"
			j := i
			num := 0
			// 后面也是 数字字符
			for j < n && unicode.IsDigit(rune(s[j])) {
				num = num*10 + int(s[j]-'0')
				j++
			}
			res += sign * num
			// 因为j 在满足之后前进了一步，把i 后移，处理这个非数字的字符
			i = j - 1
		case '(':
			// 进入括号，把之前的结果入栈，计算括号内的
			stack = append(stack, res)
			// 括号之前的符号也要入栈
			stack = append(stack, sign)
			// 因为要计算括号内的值，把res 和sign 重置
			res = 0
			sign = 1
		case ')':
			// 把括号内的值和括号外的值计算出来
			// （ 之前的计算值
			prevRes := stack[len(stack)-2]
			// ( 之前的符号
			sign := stack[len(stack)-1]
			// 出栈 原先的value
			stack = stack[:len(stack)-2]
			res = prevRes + sign*res
		case '+':
			sign = 1
		case '-':
			sign = -1
		}
	}
	return res
}
