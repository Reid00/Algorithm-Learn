package dptwodim

// canPartition 分割等和子集
// 给你一个 只包含正整数 的 非空 数组 nums 。
// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
// 0-1 背包问题
// sum(nums)/2 作为背包的容量
// nums[i] 是物品的重量和价值
func canPartition(nums []int) bool {
	// 	输入：nums = [1,5,11,5]
	// 输出：true
	// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
	n := len(nums)
	weight := 0
	for _, v := range nums {
		weight += v
	}
	// 总和是奇数 不能等分割
	if weight&1 == 1 {
		return false
	}

	weight = weight / 2

	dp := make([]int, weight+1)

	for i := 0; i < n; i++ { // 遍历物品

		for j := weight; j >= nums[i]; j-- { // 遍历容量
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			// 此处优化，提前返回
			if dp[j] == weight {
				return true
			}
		}
	}

	return dp[weight] == weight
}
