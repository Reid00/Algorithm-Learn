package math

// mySqrt x 的平方根
// 二分法
func mySqrt(x int) int {

	l, r := 0, x

	for l <= r {
		mid := l + (r-l)>>1

		if mid*mid == x {
			return mid
		}

		if mid*mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l*l > x {
		return l - 1
	}
	return l
}
