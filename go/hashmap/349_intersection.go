package hashmap

// intersection 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。
// 输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
func intersection(nums1 []int, nums2 []int) []int {
	// hashmap 记录某个数组中的元素，然后和另一个数组进行比较
	hmap := make(map[int]struct{})

	for _, v := range nums1 {
		hmap[v] = struct{}{}
	}

	hmapResult := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := hmap[v]; ok {
			if _, ok := hmapResult[v]; !ok {
				hmapResult[v] = struct{}{}
			}
		}
	}

	result := make([]int, 0)
	for k := range hmapResult {
		result = append(result, k)
	}
	return result
}
