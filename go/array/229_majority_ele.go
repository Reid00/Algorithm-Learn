package array

func majorityElement_1(nums []int) []int {
	// hash map 忽略 不写

	// 摩尔投票法
	// 1. 第一遍遍历选出可能是超过k/n(本题中1/3) 的元素
	// 2. 第二遍遍历计数确实是否大于1/3
	var element1, vote1 int
	var element2, vote2 int

	for _, v := range nums {

		if vote1 > 0 && v == element1 { // 第一个元素
			vote1++
		} else if vote2 > 0 && v == element2 { // 第二个元素
			vote2++
		} else if vote1 == 0 {
			element1 = v
			vote1++
		} else if vote2 == 0 {
			element2 = v
			vote2++
		} else {
			vote1--
			vote2--
		}
	}

	var cnt1, cnt2 int
	// 第二遍遍历确定 投票出来的元素是超过1/3个还是随机下来的元素
	for _, v := range nums {
		if vote1 > 0 && v == element1 {
			cnt1++
		}
		if vote2 > 0 && v == element2 {
			cnt2++
		}
	}

	result := make([]int, 0)
	if vote1 > 0 && cnt1 > len(nums)/3 {
		result = append(result, element1)
	}

	if vote2 > 0 && cnt2 > len(nums)/3 {
		result = append(result, element2)
	}
	return result
}
