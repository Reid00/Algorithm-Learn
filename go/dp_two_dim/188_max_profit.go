package dptwodim

// maxProfit 买卖股票的最佳时机
// dp: 你最多可以完成 k笔 交易
// Time: O(n), Space: O(n)
func maxProfit4(k int, prices []int) int {
	// dp[i][j] 表示第i天，处于j状态时的最大利润,
	// 因为可以买卖两次，j 有k种状态
	// j=1 第一次持有
	// j=2 第一次不持有
	// j=3 第二次持有
	// j=4 第二次不持有
	// ....
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

	// 有k次交易，则有2k 个状态，由于状态0 不初始化，即从1..2+1个状态
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*k+1)
	}

	// 初始化第0（从0计数）天不同状态j的最大利润
	for j := 1; j < 2*k+1; j++ {
		// 奇数表示第j次持有
		if j&1 == 1 {
			dp[0][j] = -prices[0]
		} else {
			dp[0][j] = 0
		}
	}

	// 状态转移方程赋值
	for i := 1; i < n; i++ {
		for j := 1; j < 2*k+1; j++ {
			// j 是是奇数表示持有时的最大利润
			if j&1 == 1 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	return dp[n-1][2*k]
}
