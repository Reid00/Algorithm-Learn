package dptwodim

import "math"

// maxProfit 买卖股票的最佳时机
// 动态规划，只能一次买卖
// Time: O(n), Space: O(n)
func maxProfit(prices []int) int {
	// dp[i][j]	表示第i天，现在有两种状态, j=0,表示持有股票; j=1，表示不持有股票
	// so dp[i][0] 表示第i天，持有股票能获得的最大利润
	// dp[i][1] 表示第i天，不持有股票能获取的最大利润
	// 递推公式
	// dp[i][0] 第i天 持有股票，有两种情况，第i-1天，如果持有股票，因为只能一次买卖，所以dp[i][0] = dp[i-1][0]
	// 如果第i-1 天，没有持有股票，dp[i][0] = - prices[i]
	// 即: dp[i][0] = max(dp[i-1][0], -prices[i])
	// dp[i][1] 第i天，不持有股票，有两种情况，第i-1天，本身就不持有股票，则dp[i][1]= dp[i-1][1]
	// 如果第i-1 天，持有股票，那第i天就需要进行卖掉，dp[i][1] = dp[i-1][0] + prices[i]
	// 即: dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])

	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// 初始化dp
	dp[0][0] = -prices[0] // 第0天持有股票的最大利润
	dp[0][1] = 0          // 第0天不持有股票的最大利润,1. 当天买卖是0， 2. 当天没有买卖也是0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	// 不持有股票比持有股票利润更高
	return dp[n-1][1]
}

// maxProfit 买卖股票的最佳时机
// 贪心算法: 因为股票就买卖一次，可以寻找最低点和最高点进行买卖
// Time: O(n), Space: O(1)
func maxProfit1(prices []int) int {
	low, res := math.MaxInt32, 0

	for _, v := range prices {
		low = min(low, v)
		res = max(res, v-low)
	}
	return res
}
