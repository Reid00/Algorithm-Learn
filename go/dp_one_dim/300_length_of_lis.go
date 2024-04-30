package dponedim

// lengthOfLIS 最长递增子序列
func lengthOfLIS(nums []int) int {

	// dp[i] 代表以i为结尾的递增子序列长度 (包含i)
	// 我们的状态转移方程：
	// dp[i] 的大小应该在nums[i] > nums[j] 的时候，0<=j<i； dp[i] = dp[j] + 1
	// 因为j是的大小是[0, i-1] 的，所以dp[i] 每次都有更新 所以如下
	// dp[i]=max(dp[i],dp[j]+1),0≤j<i,nums[j]<nums[i]

	n := len(nums)
	// 定义dp
	dp := make([]int, n)

	// 初始化, 最小的子序列等于nums[i] 本身 即1
	for i := 0; i < n; i++ {
		// 默认的递增子序列长度是1
		dp[i] = 1
	}

	// 状态转移方程
	res := 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				res = max(res, dp[i])
			}
		}
	}

	return res
}
