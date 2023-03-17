package array

import "sort"

func majorityElement(nums []int) int {
	// Hash 记录每个元素出现的次数，如果某个元素出现次数大于n/2 则为expected
	// 时间复杂度O(n) + O(n) = 2 O(n), 空间复杂度O(n) 不是O(1)
	var hmap = make(map[int]int)

	for _, v := range nums {
		hmap[v] += 1
	}

	for k, v := range hmap {
		if v > len(nums)>>1 {
			return k
		}
	}
	panic("no majority elements")
}

func majorityElement2(nums []int) int {
	// 排序，众数一定位于n/2 的位置
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 摩尔投票法（Boyer–Moore majority vote algorithm) 如何在任意多的候选人中（选票无序），选出获得票数最多的那个。
// 算法可以分为两个阶段：
// 对抗阶段：分属两个候选人的票数进行两两对抗抵消
// 计数阶段：计算对抗结果中最后留下的候选人票数是否有效
func majorityElement3(nums []int) int {
	// major 被计票的元素, count 票数
	var major, count int

	// 	当我们遍历下一个选票时，判断当前 count 是否为零：
	//  若 count == 0，代表当前 major 空缺，直接将当前候选人赋值给 major，并令 cousort
	//  若 count != 0，代表当前 major 的票数未被完全抵消，因此令 count--，
	//  即使用当前候选人的票数抵消 major 的票数

	for _, v := range nums {
		if count == 0 {
			major = v
		}
		if major != v {
			count--
		} else {
			count++
		}
	}
	return major
}
