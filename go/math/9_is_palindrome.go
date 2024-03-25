package math

// isPalindrome 回文数  121 yes, 123 no
func isPalindrome(x int) bool {
	// METHOD 1: 转化成字符串，双指针遍历
	// METHOD 2: 将x 的每个数反转得到的新数和x相同则是

	if x < 0 {
		return false
	}

	res := 0
	carry := 0
	tmp := x

	for x > 0 {
		carry = x % 10
		res = res*10 + carry
		x /= 10
	}
	return res == tmp
}
