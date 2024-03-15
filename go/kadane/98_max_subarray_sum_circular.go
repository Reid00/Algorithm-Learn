package kadane

import "fmt"

// maxSubarraySumCircular 环形子数组的最大和
func maxSubarraySumCircular(nums []int) int {
	// kadane 算法
	// 最大子数组 成环 --> 看下图，总和减去不成环的子数组最小和(total - minSum) === 成环区域的最大和
	total := 0
	for _, v := range nums {
		total += v
	}
	// 包含nums[i] 到i为止的元素和
	maxEndingHere := nums[0]
	minEndingHere := nums[0]
	// 全局最大元素和
	maxSoFar := nums[0]
	minSoFar := nums[0]

	for i := 1; i < len(nums); i++ {

		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		minEndingHere = min(minEndingHere+nums[i], nums[i])

		maxSoFar = max(maxSoFar, maxEndingHere)
		minSoFar = min(minSoFar, minEndingHere)
	}

	// NOTE 由于算法使用了 Kadane's algorithm，它假定输入数组中至少有一个正数元素。
	// 否则，最大和将始终为负数，导致算法返回错误的结果。因此，如果输入数组中没有正数元素，
	// 则应该特殊处理此情况，例如返回数组中的最大值（即所有元素中的最大负数）。
	// [-1,-2,-9]
	fmt.Println("min, max: ", minSoFar, maxSoFar)
	if (total - minSoFar) > 0 {
		return max(maxSoFar, total-minSoFar)
	}

	return maxSoFar
}
