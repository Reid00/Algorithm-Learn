package dptwodim

// numSquares 完全平方数
// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方
// 其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
func numSquares(n int) int {
	// 返回最小数量，用完全背包的视角
	// dp[j] 表示要实现j的完全平方数的最小数量是dp[j]
	// dp[j] = min(dp[j], dp[j-i*i] + 1)
	// 1..n 表示物品的重量和value
	// n 表示背包的最大容量

	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	for i := 0; i*i <= n; i++ { // 遍历物品
		for j := i * i; j <= n; j++ { // 遍历背包的容量
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}

	if dp[n] == n+1 {
		return -1
	}

	return dp[n]
}
