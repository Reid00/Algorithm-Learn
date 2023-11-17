package array

// removeDuplicates2 删除有序数组中的重复项 II
func removeDuplicates2(nums []int) int {

	length := len(nums)

	if length <= 2 {
		return length
	}

	// 双指针

	slow, fast := 2, 2

	for fast < length {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
