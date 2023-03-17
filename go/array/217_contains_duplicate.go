package array

func containsDuplicate(nums []int) bool {
	// HashMap

	hmap := make(map[int]struct{}) // {val: struct{}}

	for _, v := range nums {
		if _, ok := hmap[v]; ok {
			return true
		}
		hmap[v] = struct{}{}
	}
	return false
}
