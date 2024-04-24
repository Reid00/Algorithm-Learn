package dptwodim

// maxProfit 买卖股票的最佳时机
// dp: 不限买卖次数
// Time: O(n), Space: O(n)
func maxProfit2(prices []int) int {
	// dp[i][j] 表示第i天在持有(j=0) 和 不持有(j=1)股票的时候，最大利润
	// dp[i][0] 表示第i天持有股票的最大利润，如果第i-1天，持有股票，则dp[i][0] = dp[i-1][0]
	// 如果 i-1天，不持有股票，则第i天持有的利润是i-1天不持有的利润-prices[i]，即dp[i][0] = dp[i-1][1] - prices[i]
	// dp[i][1] 表示第i天不持有股票的最大利润，如果第i-1天，持有股票，则第i天需要卖出，则dp[i][1] = dp[i-1][0] + prices[i]
	// 如果i-1天 不持有股票，则第i天不变，dp[i][1] = dp[i-1][1]

	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// 初始化
	dp[0][0] = -prices[0] //第0天，持有股票的最大利润
	dp[0][1] = 0          // 第0天，不持有股票的最大利润

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
}

// maxProfit 买卖股票的最佳时机
// 贪心算法: 不限买卖次数
// Time: O(n), Space: O(1)
func maxProfit_(prices []int) int {
	// 贪心算法，只要有利润就卖

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
