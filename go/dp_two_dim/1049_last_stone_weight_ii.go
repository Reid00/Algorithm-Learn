package dptwodim

// lastStoneWeightII 最后一块石头的重量 II
func lastStoneWeightII(stones []int) int {
	// 因为提示中给出1 <= stones.length <= 30，1 <= stones[i] <= 1000，所以最大重量就是30 * 1000 。
	// 尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了。
	// stones[i] 表示物品的重量， 同时也是物品的价值
	// dp[j] 表示背包容量为j 的情况下，能够容纳的最大重量是dp[j]

	sum := 0
	for i := 0; i < len(stones); i++ {
		sum += stones[i]
	}
	// 背包的最大值是30 * 1000
	// 因为分成两堆，只需要15 *1000 的大小即可
	bagSize := sum / 2
	dp := make([]int, bagSize+1)
	dp[0] = 0

	for i := 0; i < len(stones); i++ {
		for j := bagSize; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	// 剩余背包最大的重量是
	return sum - dp[bagSize]*2
}
