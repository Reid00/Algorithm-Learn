package dponedim

// lengthOfLIS 最长递增子序列
func lengthOfLIS(nums []int) int {

	// dp[i] 代表以i为结尾的递增子序列长度 (包含i)
	// 我们的状态转移方程为：dp[i]=max(dp[i],dp[j]+1),0≤j<i,nums[j]<nums[i]

	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		// 默认的递增子序列长度是1
		dp[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 遍历dp表，返回最长的
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}
