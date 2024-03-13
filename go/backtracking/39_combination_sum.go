package backtracking

// combinationSum 组合总和
func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)

	// backtrack 回溯，start_idx 为了避免重复剪枝 比如[3,4,5],target=9 中[4,5]与[5,4]是一样的
	var backtrack func(startIdx int, total int)
	backtrack = func(startIdx, total int) {

		if total > target {
			return
		}

		if total == target {
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
			return
		}

		// 剪枝不选择之前选择过的元素了
		for i := startIdx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			// 递归， 因为同一个元素可以重复选择，所以startIdx 为i
			backtrack(i, total+candidates[i])
			// 回溯
			path = path[:len(path)-1]
		}

	}

	backtrack(0, 0)
	return res
}
