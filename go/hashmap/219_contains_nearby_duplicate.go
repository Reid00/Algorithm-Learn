package hashmap

// containsNearbyDuplicate   存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	// HashMap key: nums的元素， val: 每个元素对应的下标

	hmap := make(map[int][]int)

	for idx, v := range nums {
		hmap[v] = append(hmap[v], idx)
	}

	for _, v := range hmap {
		if len(v) <= 1 {
			continue
		}

		// 判断v []int 是不是存在元素之差小于等于k
		for i := 1; i < len(v); i++ {
			if v[i]-v[i-1] <= k {
				return true
			}
		}
	}
	return false
}

// 优化HashMap 方法，只需要一次遍历
// containsNearbyDuplicate   存在重复元素 II
func containsNearbyDuplicate2(nums []int, k int) bool {
	// HashMap key: nums的元素，
	// val: 最新该元素的下标，因为计算距离的最小值，只需要知道最相近的元素下标即可
	hmap := make(map[int]int)

	for idx, v := range nums {
		if i, ok := hmap[v]; ok && (idx-i) <= k {
			return true
		}
		hmap[v] = idx
	}
	return false
}

func containsNearbyDuplicate3(nums []int, k int) bool {
	// 滑动窗口的解法，在长度为k 的窗口中是否存在两个相同的元素
	set := make(map[int]struct{})

	for i, v := range nums {
		// 窗口set 中的元素超过k个，删除最左边的元素
		if i > k {
			delete(set, nums[i-k-1])
		}

		// 窗口中存在元素v
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
