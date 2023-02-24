package main

func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// 双指针
	var left, right int

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}

}
