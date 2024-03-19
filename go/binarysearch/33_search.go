package binarysearch

// search 搜索旋转排序数组: 两部分组成，每个部分都是单调递增的。并且第一部分的最小值大于第二部分的最大值
// 思路: 二分查找，如果mid == target 直接返回
// 如果不相等，判断mid 处于哪一部分的数组，
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.
func search(nums []int, target int) int {

	n := len(nums)

	l, r := 0, n-1

	for l <= r {

		mid := l + (r-l)>>1

		if nums[mid] == target {
			return mid
		}

		// 判定 [l:mid] 是有序数组，[mid+1:r] 是旋转数组
		if nums[l] <= nums[mid] {
			// 因为有序，则可以快速判读target 是否在有序列表中
			// NOTE nums[l] 要<=
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右侧是有序列表，判断target 是否在有序列表中
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}

	}
	return -1
}
