package binarysearch

// singleNonDuplicate 有序数组中的单一元素
// O(log(n)) 二分
// 如果不存在一个单一元素，那么nums 中所有相同的两个元素，第一位都是偶数(0,2,4)
// 因为存在了单一元素x, 导致x 的左侧满足第一位是偶数，但是右侧相反，第一位都是奇数了
// 利用上面这一点做二分查找，即 如果 mid 所在的 [偶数下标, 奇数下标] 的值相等，说明前面半段没有缺失的数
func singleNonDuplicate(nums []int) int {

	n := len(nums)

	l, r := 0, n-1

	for l < r {

		mid := l + (r-l)>>1
		// 不用区分mid 是 偶数或者奇数的写法
		// 因为偶数与1 异或相当于取反，+1, 奇数与1异或相当于-1

		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}

		// { // 如果确认 mid是偶数，他的成对的元素为奇数idx 是mid + 1
		// 	if mid&1 == 0 {

		// 		// 说明单一元素 x 在mid的右侧
		// 		if nums[mid] == nums[mid+1] {
		// 			l = mid + 2
		// 		} else {
		// 			// BUG mid 可能刚好为单元素x
		// 			// r = mid - 1
		// 			r = mid
		// 		}
		// 	} else {
		// 		// 如果确认mid 是奇数，他的成对的元素的idx 是mid-1

		// 		// 说明单一元素 x 在mid 的右侧
		// 		if nums[mid] == nums[mid-1] {
		// 			l = mid + 1
		// 		} else {
		// 			r = mid
		// 		}
		// 	}
		// }

	}
	return l
}

// singleNonDuplicate 有序数组中的单一元素
// O(n) 异或
func singleNonDuplicate2(nums []int) int {

	ans := 0

	for _, v := range nums {
		ans ^= v
	}
	return ans
}
