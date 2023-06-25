package binarysearch

// mySqrt 求x 的算术平方根整数部分
// 即要求 n * n = x , 或者 n * n < x < (n+1) * (n+1) 可以去n
func mySqrt(x int) int {

	left, right := 0, x

	for left <= right {

		mid := left + (right-left)>>1

		square := mid * mid

		if square == x {
			return mid
		}

		if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	if left*left > x {
		return left - 1
	}
	return left
}
