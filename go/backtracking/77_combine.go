package backtracking

import "slices"

// combine 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {

	// 保存所有结果的集合
	res := make([][]int, 0)
	// 满足条件的路径
	path := make([]int, 0, k)

	// 回溯算法， startIdx 代表从n 中开始取的数字
	var backtrack func(int)
	backtrack = func(startIdx int) {

		if len(path) == k {
			p := slices.Clone(path)
			res = append(res, p)
			return
		}

		// 优化剪枝
		// i 选择之后，剩余的元素个数应该至少k个，才有遍历的意义，比如n=4,k=4 那至少循环开始之后有4个元素可选择
		// 所以 i <= n - (k - len(path)) + 1
		// for i := startIdx; i < n+1; i++ {
		for i := startIdx; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)

			backtrack(i + 1)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	backtrack(1)

	return res
}
