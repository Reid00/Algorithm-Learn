package binarysearch

// findMin 寻找旋转排序数组中的最小值
func findMin(nums []int) int {

	n := len(nums)

	l, r := 0, n-1

	for l <= r {

		mid := l + (r-l)>>1

		// 如果旋转次数小于n, 那么左半段有序列表的第一个值，大于右边的所有元素
		// 如果旋转次数等于n， 整体是有序列表

		// 抛弃右侧的数组数据
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 因为上述循环条件使用 l <= r 的时候，最后 l = mid = r 的时候，会执行 l = mid + 1
	// 所以此时返回 nums[l-1] 或者 nums[r], 举例[*, 10, 1, *], nums[l]=10, nums[r]=1
	return nums[r]
}
