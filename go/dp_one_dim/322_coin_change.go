package dponedim

import "math"

// coinChange 零钱兑换
func coinChange(coins []int, amount int) int {
	// coins [1,2,5] aoumt 11
	// dp[11], 取1 dp[10] + 1
	// dp[11], 取2 dp[9] + 1
	// dp[11], 取5 dp[6] + 1
	// dp[11] = min(dp[10], dp[9], dp[6]) +1
	// dp[n] = min(amount - dp[i]) + 1

	dp := make([]int, amount+1)
	// 初始化位最大值
	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	// 穷举每一个目标数的最小硬币数
	for i := 1; i <= amount; i++ {
		//对每种子问题都穷举
		for _, v := range coins {
			if v <= i {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}
