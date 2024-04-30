package dponedim

// rob 打家劫舍
// 优化空间复杂度
func rob(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	// 使用常数空间代替O(n) 的空间
	prev := 0
	cur := 0
	for _, v := range nums {
		// 循环开始是 cur 表示dp[k-1], prev 表示dp[k-2]
		prev, cur = cur, max(cur, v+prev)
	}
	return cur
}

// 优化前
func rob_(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] 表示考虑下标i（包括i）以内的房屋，最多可以偷窃的金额为dp[i]。
	// 如果偷i房屋，说明i-1 是不可以偷的，则dp[i] = dp[i-2] + nums[i]
	// 如果不偷i房屋，则i 和 i-1 的结果是相同的； 即dp[i] = dp[i-1]
	// dp[i] = max(dp[i-1], dp[i-2] + nums[i])

	n := len(nums)
	dp := make([]int, n)

	// 初始化
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	}

	return dp[n-1]
}
