package dponedim

// change 零钱兑换II
// 计算并返回可以凑成总金额的硬币组合数
func change(amount int, coins []int) int {

	// dp[i] = dp[i] + dp[i-coin]

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {

		// 穷举所有amout 的的组合数
		for i := 1; i < amount+1; i++ {
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[amount]
}
