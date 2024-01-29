package array

// insert 插入区间
func insert(intervals [][]int, newInterval []int) [][]int {

	// 1. 遍历intervals 元素v 如果v[1] < newInterval[0], 无交集 在要插入位置的左侧，直接加入新切片中
	// 2. if v[0] > newInterval[1], 无交集，在要插入位置的右侧
	// 2. 其他情况，和要插入位置的元素{left, right} 有交集，比较和left，right 的大小，进行merge

	result := make([][]int, 0)
	left, right := newInterval[0], newInterval[1]

	// 初始不需要merge
	merged := false

	for _, v := range intervals {

		if v[1] < left {
			result = append(result, v)
		} else if v[0] > right {

			// v 在要插入位置i的右侧，所以先插入一个元素在标识
			// merged 为false 说明i 处没有元素，临时插入，把merge 设为false
			if !merged {
				result = append(result, []int{left, right})
				merged = true
			}
			// 再插入右侧元素
			result = append(result, v)
		} else {
			// 其他情况有交集，则和位置i的元素比较和并 得到最终的元素
			left = min(v[0], left)
			right = max(v[1], right)
		}
	}

	// 循环结束后 如果merge为false 说明还没插入元素, 可能需要插入在最后
	if !merged {
		result = append(result, []int{left, right})
	}

	return result
}
