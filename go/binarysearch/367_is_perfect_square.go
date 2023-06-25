package binarysearch

// isPerfectSquare 有效的完全平方数
// 0..num 二分查找
func isPerfectSquare(num int) bool {

	left, right := 0, num

	for left <= right {
		mid := left + (right-left)>>1

		square := mid * mid

		if square > num {
			right = mid - 1
		} else if square < num {
			left = mid + 1
		} else {
			return true
		}
	}

	return false

}
