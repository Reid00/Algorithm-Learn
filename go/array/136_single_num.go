package array

func singleNumber(nums []int) int {
	// 位运算 ^ 相同为0，不同为1
	// 0 和任何数 异或都等于 任何数自己
	var i = 0

	for _, v := range nums {
		i ^= v
	}
	return i
}
