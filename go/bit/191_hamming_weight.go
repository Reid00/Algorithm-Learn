package bit

// hammingWeight 位 1 的个数
func hammingWeight(num uint32) int {
	// x&(x-1) 将二进制中 最右边的1 改为 0
	var res int

	for num != 0 {
		num = num & (num - 1)
		res++
	}

	return res
}
