package array

func containsNearbyDuplicate(nums []int, k int) bool {
	// HashMap
	hmap := make(map[int]int) // {nums[i]: i}

	for i, v := range nums {
		// 可能存在多个duplicated 的ele
		if val, ok := hmap[v]; ok {
			if i-val <= k {
				return true
			}
		}
		hmap[v] = i
	}

	return false

}

func containsNearbyDuplicate2(nums []int, k int) bool {
	// sliding window
	set := make(map[int]struct{})

	for i, v := range nums {

		if i > k {
			delete(set, nums[i-k-1])
		}

		if _, ok := set[v]; ok {
			return true
		}

		set[v] = struct{}{}
	}
	return false
}
