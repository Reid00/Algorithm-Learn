package dptwodim

// maxProfit 买卖股票的最佳时机
// dp: 你最多可以完成 两笔 交易
// Time: O(n), Space: O(n)
func maxProfit3(prices []int) int {
	// dp[i][j] 表示第i天，处于j状态时的最大利润,
	// 因为可以买卖两次，j 有四种状态
	// j=1 第一次持有
	// j=2 第一次不持有
	// j=3 第二次持有
	// j=4 第二次不持有
	// dp[i][1] 表示第i天，第一次持有获取的最大利润, 如果i-1 已经持有,dp[i][1] = dp[i-1][1]
	// 如果i-1 没有持有，dp[i][1] = -prices[i]
	// dp[i][2] 表示第i天，第一次不持有获取的最大利润，如果i-1 已经持有，需要卖出,dp[i][2] = dp[i-1][1] + prices[i]
	// 如果i-1 没有持有，dp[i][2] = dp[i-1][2]
	// dp[i][3] 表示第i天，第二次持有获取的最大利润, 如果i-1 已经持有，dp[i][3] = dp[i-1][3]
	// 如果i-1 没有持有，dp[i][3] = dp[i-1][2] - prices[i]
	// dp[i][4] 表示第i天，第二次不持有获取的最大利润，if i-1 已经持有，dp[i][4] = dp[i-1][3] + prices[i]
	// 如果i-1 没有持有, dp[i][4]= dp[i-1][4]

	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0] // 第一次持有的最大利润
	dp[0][2] = 0          // 第一次不持有，当天买，当天卖 则是0
	dp[0][3] = -prices[0] //第二次持有，买卖之后利润是0， 再买入
	dp[0][4] = 0

	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], -prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	// 第二次卖出的利润最大，可以理解为第一次卖出利润最大的话，也可以立即进行第二买卖
	return dp[n-1][4]
}
