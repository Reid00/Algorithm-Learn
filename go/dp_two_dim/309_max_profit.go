package dptwodim

// 买卖股票的最佳时机含冷冻期 1天
// dp: 无交易次数限定
// Time: O(n), Space: O(n)
func maxProfit5(prices []int) int {
	// dp[i][j] 表示第i天， 处于状态j 的时候最大利润
	// 由于含有冷冻期，状态j 有以下四种状态
	// 1. 持有股票
	// 2. 不持有股票 分为两种:
	// 2-1 保持卖出股票的状态(前两天已经卖出股票了，度过了一天(i-1)冷冻期，或者更早一天是卖出状态，没操作)
	// 2-2 今天卖出股票
	// 3. 冷冻期
	// 递推公式推导:
	// status 0, dp[i][0] 表示第i天持有股票，如果i-1天持有股票, 则dp[i][0] = dp[i-1][0]
	// if i-1天不持有股票, 并且为冷冻期, dp[i][0] = dp[i-1][3] - prices[i]
	// if i-1天不持有股票，不为冷冻期（只能为状态2-1）, dp[i][0] = dp[i-1][1] - prices[i]
	// status 1, dp[i][1] 表示第i天不持有股票，保持这个状态；if i-1 是这个状态 dp[i][1] = dp[i-1][1]
	// if i-1 是冷冻期， dp[i][1] = dp[i-1][3]
	// status 2, dp[i][2] 表示第i天卖出股票，则i-1天一定是持有股票, dp[i][2] = dp[i-1][0] + prices[i]
	// status 3, dp[i][3] 表示第i天是冷冻期，则i-1天刚卖出股票, dp[i][3] = dp[i-1][2]

	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
	}

	// 初始化
	dp[0][0] = -prices[0] // 第一次持有股票
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[n-1][1], dp[n-1][2], dp[n-1][3])
}
