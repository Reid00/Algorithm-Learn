package binarysearch

// searchInsert 搜索插入位置
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func searchInsert(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)>>1

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 为什么返回l 是因为，最后如果target 不在 nums 中，最后一次循环l==r, 此时没有target， l = mid + 1 了
	// nums = [1,3,5,6], target = 2
	// return 1
	return l
}
