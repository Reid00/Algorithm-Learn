package dponedim

// lenLongestFibSubseq 最长的斐波那契子序列的长度
func lenLongestFibSubseq(arr []int) int {
	// 动态规划  + 哈希表
	// dp[i][j] 表示 在arr 中 以i,j 为结尾的斐波那契子序列的长度，即 i + j = k 存在
	// 默认初始长度为2, 如果最终长度ans>=3 则存在，否则返回0
	// 状态转移方程: 如果存在 arr[i] + arr[j] = arr[k]
	// 则: dp[j][k] = max(dp[j][k], dp[i][j] + 1)

	n := len(arr)
	ans := 0

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 2
		}
	}

	// {arr_val: arr_idx}
	hmap := make(map[int]int)
	for i, v := range arr {
		hmap[v] = i
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			val := arr[i] + arr[j]
			// 下一个斐波那契数存在
			if idx, ok := hmap[val]; ok {
				dp[j][idx] = max(dp[j][idx], dp[i][j]+1)
				ans = max(ans, dp[j][idx])
			}
		}
	}

	if ans >= 3 {
		return ans
	}

	return 0
}
