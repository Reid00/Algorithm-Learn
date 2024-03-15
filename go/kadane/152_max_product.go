package kadane

// maxProduct 乘积最大子数组
func maxProduct(nums []int) int {

	// Kadane
	maxEndingHere, minEndingHere := nums[0], nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {

		oldMax := maxEndingHere
		// 存在 nums[i] 是个负数，相乘后，由最大值变为最小值，最小值变为最大值
		maxEndingHere = max(maxEndingHere*nums[i], nums[i], minEndingHere*nums[i])
		minEndingHere = min(minEndingHere*nums[i], nums[i], oldMax*nums[i])

		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}
