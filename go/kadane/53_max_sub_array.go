package kadane

import "math"

// maxSubArray 最大子数组和
func maxSubArray(nums []int) int {
	// Kadane 算法
	if len(nums) == 0 {
		return 0
	}

	// maxEndingHere  表示到当前位置为止，包括当前位置在内的子数组中的最大值
	// 在遍历过程中，对于位置i，可以分两种情况考虑：
	// 一种是将当前位置加入到前面的最大子数组中，
	// 另一种是以当前位置为起点开始新的最大子数组。

	// maxSoFar 表示到当前位置为止，已经遍历过的子数组中的最大值
	maxEndingHere, maxSoFar := nums[0], nums[0]

	for i, v := range nums {
		if i == 0 {
			continue
		}
		maxEndingHere = max(maxEndingHere+v, v)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// maxSubArray 最大子数组和
func maxSubArray2(nums []int) int {
	// 前缀和 => pre_sum, pre_sum[0] = 0
	// pre_sum[i+1] = pre_sum[i] + nums[i]
	// pre_sum[i] 表示 nums[0..i-1] 的子数组和
	// nums[i..j] 的字数组和 => pre_sum[j+1] - pre_sum[j]
	// 前缀和
	preSum := 0
	// 最小的前缀和
	minPreSum := 0
	ans := math.MinInt32

	for _, v := range nums {
		// 求出到v 为止的前缀和
		preSum += v
		// 计算每次和最小前缀和的差值 => 前缀和的差值即是字数组的和
		ans = max(ans, preSum-minPreSum)
		// update
		minPreSum = min(minPreSum, preSum)
	}
	return ans
}
