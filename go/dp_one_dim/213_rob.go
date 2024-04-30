package dponedim

// rob 打家劫舍 II
// 第一个房屋和最后一个房屋是紧挨着的
func rob2(nums []int) int {

	// 因为第一个房间和最后一个房间首位相连 so 如果偷窃了第一间房屋，则不能偷窃最后一间房屋
	// 因此分两种范围偷第一个: [0, n-2] 和 不偷第一个：[1, n-1]

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	res1 := robRange(nums, 0, n-2)
	res2 := robRange(nums, 1, n-1)
	return max(res1, res2)
}

// robRange 可以偷房子的范围 [0, n-2] or [1, i-1]
// end including
func robRange(nums []int, start, end int) int {

	dp := make([]int, len(nums))

	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[end]
}
