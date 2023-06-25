package binarysearch

// searchRange 在排序数组(非递减)中查找元素的第一个和最后一个位置
// 二分
func searchRange(nums []int, target int) []int {
	n := len(nums)

	l, r := 0, n-1
	ans := make([]int, 0)

	for l <= r {
		mid := l + (r-l)>>1

		if nums[mid] == target {

			// 往左查看是否还有等于target 的值, 开始位置
			left := mid
			for left >= 0 && nums[left] == target {
				left--
			}
			// 退出上述循环时，left -1 了
			ans = append(ans, left+1)

			// 结束位置的坐标
			right := mid
			for right < n && nums[right] == target {
				right++
			}
			ans = append(ans, right-1)
			return ans
		}

		// mid 大于target 因为nums 是非递减切片
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return []int{-1, -1}
}
