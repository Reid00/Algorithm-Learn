package dponedim

// maxSubArray 最大字数组和
func maxSubArray(nums []int) int {

	// dp[i] 代表以i为结尾的最大字数组合
	// dp[i] = nums[i]  if dp[i-1] < 0
	// dp[i] = dp[i-1] + nums[i] if dp[i-1] >= 0

	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {

		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}
