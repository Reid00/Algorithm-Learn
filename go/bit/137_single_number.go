package bit

// 只出现一次的数字 II
// 除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
func singleNumber2(nums []int) int {
	// 如果每个元素出现了三次，那么他的每位二进制bit位 也出现了三次
	// 所以统计每个元素的二进制位的个数和，然后对3求模，得到二进制表示转为十进制即是结果

	var res int32

	// nums 中的元素是32位数字，所以遍历32次，计算每次最后以为
	for i := 0; i < 32; i++ {
		// 初始没位置位0
		cnt := int32(0)

		for _, v := range nums {
			// >> i 位，右移i位，意味着计算到元素v 的第i位的数字了
			// &1 取出该数字
			cnt += int32(v) >> int32(i) & 1
		}
		// 将计算的cnt 左移 i 位，然后赋值到res 的i 位
		res |= cnt % 3 << i
	}

	return int(res)
}
