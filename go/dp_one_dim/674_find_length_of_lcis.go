package dponedim

// findLengthOfLCIS 最长连续递增序列
// 有别于第300题最长递增序列，此题要连续递增
func findLengthOfLCIS(nums []int) int {

	// dp[i] 表示以i 为结尾的连续递增序列的长度是dp[i]
	// if dp[i] > dp[i-1], dp[i] = dp[i-1] + 1

	n := len(nums)

	dp := make([]int, n)

	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			ans = max(ans, dp[i])
		}
	}

	return ans
}

// 贪心，局部最优就是全局最优
func findLengthOfLCIS_(nums []int) int {

	win := 1
	ans := 1

	for i, v := range nums {
		if i>0 && v > nums[i-1] {
			win++
			ans = max(ans, win)
		} else {
			win = 1
		}
	}
	return ans
}
