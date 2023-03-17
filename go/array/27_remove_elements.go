package array

func removeElement(nums []int, val int) int {

	// 快慢指针
	var slow, fast int

	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
