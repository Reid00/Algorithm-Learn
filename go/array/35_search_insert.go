package array

// searchInsert 搜索插入位置
// 请必须使用时间复杂度为 O(log n) 的算法。
// BinarySearch
func searchInsert(nums []int, target int) int {

	n := len(nums)

	l, r := 0, n-1
	ans := n

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			ans = mid
			r = mid - 1
		}
	}

	return ans

}
