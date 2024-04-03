package dptwodim

// coinChange 零钱兑换
// 返回能够兑换的最小硬币个数
// 完全背包的视角
func coinChange(coins []int, amount int) int {
	// dp[j] 表示能够凑到amount 所需的最小硬币个数
	// coins[i] 表示物品和物品的价值
	// amount 是背包的容量
	// dp[j] = min(dp[j], dp[j-nums[i]] + 1)

	dp := make([]int, amount+1)
	// 初始化
	// 因为求最小值，初始给个较大值 amount + 1
	dp[0] = 0
	// 考虑到状态转移方程的特性，初始化要是一个大的值
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	// 不能凑够
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
