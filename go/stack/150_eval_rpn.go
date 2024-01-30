package stack

import "strconv"

// evalRPN 逆波兰表达式求值
func evalRPN(tokens []string) int {
	// 栈，
	// 碰到 有效的算符为 '+'、'-'、'*' 和 '/' 。 将其前两个元素出栈进行计算
	stack := make([]int, 0)

	for _, v := range tokens {

		num, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, num)
		} else {
			// 转化出错说明是有效运算符
			left := stack[len(stack)-2]
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			var res int
			switch v {
			case "+":
				res = left + right
			case "-":
				res = left - right
			case "*":
				res = left * right
			case "/":
				res = left / right
			}
			stack = append(stack, res)
		}
	}

	return stack[0]
}
