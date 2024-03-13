package backtracking

import "slices"

// 子集
func subsets(nums []int) [][]int {

	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(startIdx int)

	backtrack = func(startIdx int) {
		p := slices.Clone(path)
		res = append(res, p)

		for i := startIdx; i < len(nums); i++ {
			path = append(path, nums[i])

			backtrack(i + 1)
			path = path[:len(path)-1]
		}

	}

	backtrack(0)
	return res
}
