package hashmap

import "sort"

// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。
// 返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
// （如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
func intersect(nums1 []int, nums2 []int) []int {
	// hashmap 记录某个数组中的元素，和出现的次数，然后和另一个数组进行比较
	hmap := make(map[int]int)

	for _, v := range nums1 {
		hmap[v] += 1
	}

	hmapResult := make(map[int]int)
	for _, v := range nums2 {
		if _, ok := hmap[v]; ok {
			hmapResult[v] += 1
		}
	}

	result := make([]int, 0)
	for k, v := range hmapResult {
		cnt := min(v, hmap[k])
		for i := 0; i < cnt; i++ {
			result = append(result, k)
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// intersect2 排序 + 双指针
func intersect2(nums1 []int, nums2 []int) []int {

	// 先排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := make([]int, 0)
	// 双指针，比较元素大小，相同加入结果中，不同则则移动较小元素的指针
	idx1, idx2 := 0, 0

	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums2[idx2] < nums1[idx1] {
			idx2++
		} else {
			result = append(result, nums1[idx1])
			idx1++
			idx2++
		}
	}
	return result
}


