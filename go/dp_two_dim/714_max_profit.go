package dptwodim

// 买卖股票的最佳时机含手续费fee
// dp: 无交易次数限定
// Time: O(n), Space: O(n)
func maxProfit6(prices []int, fee int) int {
	// dp[i][j] 表示第i天，处于状态j 时的最大利润；此处i有两个值，0 持有股票和 1 不持有股票
	// dp[i][0]：第i天持有股票获得利润的最大值； dp[i][1]：第i天不持有股票获得利润的最大值；

	n := len(prices)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee) // 卖出的时候计算手续费
	}

	return dp[n-1][1]
}

// 买卖股票的最佳时机含手续费fee
// 贪心算法: 无交易次数限定
// Time: O(n), Space: O(1)
func maxProfit_6(prices []int, fee int) int {

	res := 0
	// 不同于122题，此处贪心的核心在于买卖是需要手续费的，要保证每次买卖的利润大于手续费
	// 因此可以把手续费需先放置到买入的时候，这样下次卖出有利润即可
	minVal := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minVal {
			minVal = prices[i]
		}

		if prices[i] > minVal+fee {
			res += prices[i] - minVal - fee
			// 重要，更新minVal, 持有股票，使其 预包含了手续费
			minVal = prices[i] - fee
		}
	}

	return res
}
