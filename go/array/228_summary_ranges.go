package array

import "strconv"

// summaryRanges 汇总区间
func summaryRanges(nums []int) []string {

	result := make([]string, 0)

	n := len(nums)

	for i := 0; i < n; {
		// i 当前idx， v 当前val  想要保持汇总区间，nums[i+1] == nums[i] +1 / nums[i+1]==v+1
		// 相邻元素值相差 1
		left := i
		for i = i + 1; i < n && nums[i-1]+1 == nums[i]; i++ {
			// 退出当前循环后表明 i 越界 或者 值相差不再是1, 不符合汇总区间
		}

		res := strconv.Itoa(nums[left])

		// 判断区间是否存在right 的方法是上面循环后 left 是否小于满足条件的最后一个元素idx i-1
		if left < i-1 {
			res += "->" + strconv.Itoa(nums[i-1])
		}

		result = append(result, res)
	}
	return result
}
