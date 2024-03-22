package bit

import "strconv"

// addBinary 二进制求和
func addBinary(a, b string) string {

	ans := ""

	// 进位
	carry := 0

	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			// 字符转成数字
			carry += int(a[lenA-i-1] - '0')
		}

		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		// 求余之后是各位数字
		ans = strconv.Itoa(carry%2) + ans
		// 求商是进位数字
		carry /= 2
	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}
