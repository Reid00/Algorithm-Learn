package binarysearch

import "math"

// findPeakElement 寻找峰值  峰值元素是指其值严格大于左右相邻值的元素。
// 二分
func findPeakElement(nums []int) int {

	n := len(nums)

	left, right := 0, n-1

	// 取出i 在nums 中的值
	//  方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i < 0 || i >= n {
			return math.MinInt64
		}
		return nums[i]
	}

	for left <= right {
		mid := left + (right-left)>>1

		if get(mid) > get(mid+1) && get(mid) > get(mid-1) {
			return mid
		}

		// mid 处在一个上升的陂位， 最大值在右侧
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}
