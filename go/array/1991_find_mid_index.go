package array

// findMiddleIndex 寻找数组的中心索引
func findMiddleIndex(nums []int) int {
	// 	先求得数组中所有元素之和sum；
	// 遍历数组，取当前下标左边的元素之和left_sum，同时sum减去已遍历元素，比较二者是否相等，相等则返回当前下标；
	// 遍历结束，代表没有中心索引，返回-1
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, v := range nums {
		if leftSum == sum-v {
			return i
		}

		leftSum += v
		sum -= v
	}
	return -1
}
