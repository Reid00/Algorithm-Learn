package array

func removeDuplicates(nums []int) int {
	// 双指针，快慢指针

	length := len(nums)
	if length <= 1 {
		return length
	}

	var slow, fast = 1, 1

	for fast = 1; fast < length; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
