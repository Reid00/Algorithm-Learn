package stack

// isValid 有效的括号
// 用栈
func isValid(s string) bool {

	lenth := len(s)
	if lenth&1 == 1 {
		return false
	}

	stack := make([]rune, 0)

	hmap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if _, ok := hmap[ch]; ok {
			// 不匹配
			if len(stack) == 0 || stack[len(stack)-1] != hmap[ch] {
				return false
			}

			// 匹配出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
