package backtracking

// letterCombinations 电话号码组合
func letterCombinations(digits string) []string {

	// 回溯
	// 当题目中出现 “所有组合” 这种字眼时，第一时间要想到回溯
	// 1. 明确所有选择：画出搜索过程的决策树，根据决策树来确定搜索路径。
	// 2. 明确终止条件：推敲出递归的终止条件，以及递归终止时的要执行的处理方法。
	// 3. 将决策树和终止条件翻译成代码：
	// 定义回溯函数（明确函数意义、传入参数、返回结果等）。
	// 书写回溯函数主体（给出约束条件、选择元素、递归搜索、撤销选择部分）。
	// 明确递归终止条件（给出递归终止条件，以及递归终止时的处理方法）。

	var phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var res = make([]string, 0)

	var backtrack func(string, int, string)
	backtrack = func(digits string, idx int, path string) {
		// 递归终止条件
		if idx == len(digits) {
			res = append(res, path)
			return
		}

		// 取出数字对应的字符串
		digit := digits[idx]
		letters := phoneMap[string(digit)]

		for i := 0; i < len(letters); i++ {
			// 选择元素
			letter := string(letters[i])
			backtrack(digits, idx+1, path+letter)
			// 回溯正常应该还有一个撤销操作，本题不需要，因为所有路径都符合条件
		}

	}

	if len(digits) == 0 {
		return nil
	}

	backtrack(digits, 0, "")
	return res
}
