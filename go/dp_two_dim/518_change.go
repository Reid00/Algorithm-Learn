package dptwodim

// change 零钱兑换II
// 计算并返回可以凑成总金额的硬币组合数 （硬币数量无限个）
// 完全背包视角 从新思考这个问题
func change(amount int, coins []int) int {
	// coins[i] 表示物品的重量  和 价值
	// dp[j] 表示容量为j 的背包中，能够填满的方法有dp[j]种
	// dp[j] += dp[j - conis[i]]

	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}
