package dptwodim

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬至多m (1 <= m < n)个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 完全背包问题，每次可以爬m
func climbStairs3(n, m int) int {

	// dp[j] 表示容量为j的背包，装满有dp[j] 种方法
	// dp[j] 此处表示爬到j, 有dp[j] 种方法

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i <= n; i++ { //背包重量
		for j := 1; j <= m; j++ { // 遍历物品 即每次爬的台阶数量
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}

	return dp[n]
}
