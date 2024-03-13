package backtracking

// maxUniqueSplit 拆分字符串使唯一子字符串的数目最大
func maxUniqueSplit(s string) int {
	// 回溯

	// 子串最终数目
	var res = 0
	// 搜索过程中的字符串子集，不能重复要用集合
	path := make(map[string]bool)

	// backtrack 回溯搜索各种子集的可能
	// idx 搜索开始的下标, count 子串的个数
	var backtrack func(idx, count int)
	backtrack = func(idx, count int) {

		// 剪枝, 剩余的字符个数小于 len(s)- idx
		// + 已经拆分的子集 count
		// 小于当前res 则没有必要继续递归搜索下去了
		if len(s)-idx+count <= res {
			return
		}

		if idx == len(s) {
			res = max(res, count)
			return
		}

		for i := idx; i < len(s); i++ {
			subStr := s[idx : i+1]
			if !path[subStr] {
				path[subStr] = true
				backtrack(i+1, count+1)
				path[subStr] = false
			}
		}
	}

	backtrack(0, 0)
	return res
}
