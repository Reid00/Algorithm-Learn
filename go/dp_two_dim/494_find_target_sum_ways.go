package dptwodim

import "math"

// findTargetSumWays 目标和
// 0-1 背包问题
func findTargetSumWays(nums []int, target int) int {

	// 本题要如何使表达式结果为target，
	// 既然为target，那么就一定有 所有+的元素和为left， 所有-的元素和为right。即left - right = target
	// 数组中所有元素的和为sum， 即sum =left + right
	// 即 target + sum = left + right + left - right 即 left = (target + sum)/2
	// target, sum 是固定的
	// 此时问题就是在集合nums中找出和为left的组合。

	// 这次和之前遇到的背包问题不一样了，之前都是求容量为j的背包，最多能装多少。
	// 本题则是装满有几种方法。其实这就是一个组合问题了。
	// dp[j] 表示：填满j（包括j）这么大容积的包，有dp[j]种方法
	// dp[j] += dp[j-nums[i]]

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 不是偶数
	if (sum+target)&1 == 1 {
		return 0
	}

	// nums 中找出和为 left 的组合，如果target 过大，会导致left > sum, 无解
	if math.Abs(float64(target)) > float64(sum) {
		return 0
	}

	bagWeight := (target + sum) >> 1

	dp := make([]int, bagWeight+1)
	// 初始化dp, 开始都为1
	// 因为一切都是从dp[0] 开始，如果dp[0] = 0, 想加都是0 了
	dp[0] = 1

	for i := 0; i < len(nums); i++ { // 遍历物品
		for j := bagWeight; j >= nums[i]; j-- { // 遍历背包容量
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bagWeight]
}
