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
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= len(nums); i++ {
		// dp[i] 有两种方案，不偷i-1， 那么问题就变成在前 k−1 个房子中偷到最大的金额，也就是子问题 f(k−1)
		// 偷i-1, 则不能偷i-2， 结果为 dp[i-2] + nums[i-1]
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}

	return dp[len(nums)]
}
