package hashmap

// longestConsecutive 最长连续序列
func longestConsecutive(nums []int) int {
	// METHOD1: 排序后遍历
	// METHOD2: HashMap 先去重，然后找到连续序列中的最小元素x
	// 把x + 1, x+2 依次递增，观察hashmap 中是否有元素

	hmap := make(map[int]bool)

	for _, v := range nums {
		hmap[v] = true
	}

	result := 0
	// 遍历hmap 先找到连续序列中的最小元素
	// k 是nums 中的元素
	for k, _ := range hmap {

		// 如果v-1 不在hmap 里面，说明v 可以是连续序列的开始的值
		// if _, ok := hmap[k-1]; !ok {
		// 或者下面的写法
		if !hmap[k-1] {

			curNum := k
			length := 1

			for hmap[curNum+1] {
				curNum++
				length++
			}

			if result < length {
				result = length
			}
		}
	}
	return result
}
