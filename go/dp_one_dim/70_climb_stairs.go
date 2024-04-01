package dponedim

// climbStairs 爬楼梯
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 用 dp 优化
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// climbStairs 爬楼梯
// 递归 无优化，超时
func climbStairs2(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	// 假设第一次爬1/ 2， 剩下就有 climbStairs(n-1)/ climbStairs(n-2) 种了
	return climbStairs(n-1) + climbStairs(n-2)
}
