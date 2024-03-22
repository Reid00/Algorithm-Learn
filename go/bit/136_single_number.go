package bit

// singleNumber 只出现一次的数字  除了某个元素只出现一次以外，其余每个元素均出现两次
func singleNumber(nums []int) int {
	// ^ 运算，x ^ y 相同为0， 不同位1
	// 所有的元素异或之后只剩下单独的那个 0^任意值=任意值

	num := 0

	for _, v := range nums {
		num ^= v
	}

	return num
}
