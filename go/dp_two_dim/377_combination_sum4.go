package dptwodim

// combinationSum4 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
// nums 中的元素可以重复使用
// 完全背包问题
func combinationSum4(nums []int, target int) int {
	// nums[i] 代表物品的重量 和 价值
	// dp[j] 表示容量为j 的背包，可以填满有dp[j] 种方法
	// dp[j] += dp[j-nums[i]]

	dp := make([]int, target+1)
	// 初始化
	dp[0] = 1

	// 此题本质是求排列，不是求组合
	// 排列问题的话， 背包容量应该放在外层
	for i := 0; i < target+1; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
