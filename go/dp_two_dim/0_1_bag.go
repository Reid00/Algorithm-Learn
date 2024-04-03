package dptwodim

// 0-1 背包问题
// 有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i]
// 每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
// weight = [1, 3, 4]
// value = [15, 20, 30]
// bagweight = 4
// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/%E8%83%8C%E5%8C%85%E6%80%BB%E7%BB%93%E7%AF%87.md

// 背包优化
// dp + 滚动数组
func weightBag2(weight, value []int, bagWeight int) int {
	// dp[j] 表示 容量为j的背包，所背的物品价值可以最大为dp[j]
	// dp[j] = max(dp[j], dp[j-weight[i]] + val[i])
	// 此时dp[j]有两个选择，一个是取自己dp[j] 相当于 二维dp数组中的dp[i-1][j]，
	// 即不放物品i，一个是取dp[j - weight[i]] + value[i]，即放物品i，指定是取最大的，毕竟是求最大价值，

	dp := make([]int, bagWeight+1)

	// 初始化，背包容量为0，的时候价值都为0
	// dp 上面声明的时候，已经初始化

	n := len(weight)

	for i := 0; i < n; i++ { //遍历物品
		// 需要注意的是，j 一定要从大到小，防止物品0 被重复计算
		for j := bagWeight; j >= weight[i]; j-- { //遍历背包容量
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}

	}
	return dp[bagWeight]
}

// weightBag
func weightBag(weight, value []int, bagWeight int) int {
	// dp[i][j] 二维，第一维i 表示选择的物品，从weight 中找
	// 第二维 j 表示背包能放下的最大重量
	// dp[i][j] 表示从下标为[0-i]的物品里任意取(包含i)，放进容量为j的背包，价值总和最大是多少
	// dp[i][j] =
	// 不放物品i：由dp[i - 1][j]推出，即背包容量为j，里面不放物品i的最大价值，此时dp[i][j]就是dp[i - 1][j]。
	// (其实就是当物品i的重量大于背包j的重量时，物品i无法放进背包中，所以背包内的价值依然和前面相同。)
	// 放物品i：由dp[i - 1][j - weight[i]]推出，dp[i - 1][j - weight[i]]
	// 为背包容量为j - weight[i]的时候不放物品i的最大价值，那么dp[i - 1][j - weight[i]] + value[i] （物品i的价值），就是背包放物品i得到的最大价值

	n := len(weight)

	dp := make([][]int, n) // 第一维物品的个数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, bagWeight+1) // 第二维，背包能放下的最大重量，包含bagWeight 本身
	}

	// 初始化dp
	// 当j为零的时候，价值最大为0
	for i := 0; i < n; i++ {
		dp[i][0] = 0
	}
	// 当i为零的时候，价值为value[0]
	// i 开始是weight[0] weight 是递增的
	for i := weight[0]; i < bagWeight+1; i++ {
		dp[0][i] = value[0]
	}

	// 为什么从1开始遍历，因为状态转移公式
	for i := 1; i < n; i++ { // 遍历物品
		for j := 0; j < bagWeight+1; j++ { // 遍历背包容量

			if j < weight[i] { // 物品不能放到背包中
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-weight[i]]+value[i])
			}
		}
	}
	return dp[n-1][bagWeight]
}
