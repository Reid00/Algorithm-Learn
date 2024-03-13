package backtracking

import "slices"

// permute 全排列
func permute(nums []int) [][]int {

	res := make([][]int, 0)

	path := make([]int, 0, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			p := make([]int, 0, len(nums))
			copy(p, path)
			res = append(res, p)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !slices.Contains(path, nums[i]) {
				path = append(path, nums[i])
				backtrack()
				path = path[:len(path)-1]
			}
		}

	}

	backtrack()
	return res
}
