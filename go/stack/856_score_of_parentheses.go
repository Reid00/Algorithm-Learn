package stack

// scoreOfParentheses 括号的分数
// Method1: 栈
func scoreOfParentheses(s string) int {

	// 括号的分数，相当于把把括号看成树，每个节点的值，等于子节点之和* 2
	// 求虚拟根节点的值

	ans := []int{0}

	for _, ch := range s {
		// 左括号压栈，并赋值初试分数为0
		if ch == '(' {
			ans = append(ans, 0)
		} else {
			// 碰到右括号(有匹配项) 出栈，出栈处的对应的左( 计算分数
			// 栈顶剩下左括号的分数 = max(2 * 出栈的(左括号内部的分数)， 1)
			
			v := ans[len(ans)-1]
			ans = ans[:len(ans)-1]
			// 所有子节点分数之和
			ans[len(ans)-1] += max(2*v, 1)
		}
	}

	return ans[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

// scoreOfParentheses2 根节点的值是 每个叶子节点值 与其深度的幂的总和。
func scoreOfParentheses2(s string) int {

	var ans, depth int

	for i, ch := range s {
		if ch == '(' {
			depth += 1
		} else {
			depth--

			if ch == ')' && s[i-1] == '(' {
				ans += 1 << depth
			}
		}
	}

	return ans
}
