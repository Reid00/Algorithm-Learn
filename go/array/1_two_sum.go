package array

// O(n) Hash Map
func twoSum1(nums []int, target int) []int {

	// hashMap 将一个num 放到hMap 中，查看另一个target - num 是不是也在其中
	// key 是num， val 是idx
	hmap := make(map[int]int, len(nums))

	for i, v := range nums {
		if idx, ok := hmap[target-v]; ok {
			return []int{i, idx}
		} else {
			hmap[v] = i
		}
	}

	return []int{}
}

// O(n2) 暴力遍历
func twoSum2(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
