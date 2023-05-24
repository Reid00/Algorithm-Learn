package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	result := make([]int, 0, len(nums1))
	// O(m*m)
	for i, v := range nums1 {
		for j, item := range nums2 {
			if v == item {
				for k := j + 1; k < len(nums2); k++ {
					if nums2[k] > v {
						result[i] = nums2[k]
					}
				}
			}

		}
		if result[i] == 0 {
			result[i] = -1
		}
	}
	return result
}

// nextGreaterElement2 O(m+n) 单调栈 + hashMap
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	hmap := make(map[int]int)

	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]

		// 遍历到n 的时候，一一比较 如果存在元素大于 n, 则这个元素(栈顶) 是他的next greater element
		for len(stack) > 0 && n >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			hmap[n] = stack[len(stack)-1]
		} else {
			hmap[n] = -1
		}

		stack = append(stack, n)
	}

	result := make([]int, len(nums1))
	for i, v := range nums1 {
		result[i] = hmap[v]
	}
	return result
}
