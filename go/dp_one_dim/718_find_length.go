package dponedim

// findLength 最长重复子数组(连续的)
func findLength(nums1 []int, nums2 []int) int {
	// dp[i][j] 表示以nums1[i-1] 下标 和 以nums2[j-1] 为下标的最长重复字数组的长度是dp[i][j]
	// 注意此处是（i-1,j-1）为下标，不是(i,j) 为下标，是为了方便初始化
	// if nums1[i-1] == nums2[j-1]: dp[i][j] = dp[i-1][j-1] + 1

	m, n := len(nums1), len(nums2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化，这里如果是以(i,j)为下标，则需要对每行，没列单独初始化
	// 但是我们以(i-1,j-1)为下标
	// dp[0][0] 是没有意义的都为0

	ans := 0

	// 状态转移方程
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			ans = max(ans, dp[i][j])
		}
	}

	return ans
}

// findLength_ 表示i 为下标 not i-1
func findLength_(nums1 []int, nums2 []int) int {
	// dp[i][j] 表示以nums1[i] 下标 和 以nums2[j] 为下标的最长重复字数组的长度是dp[i][j]
	// if nums1[i] == nums2[j]: dp[i][j] = dp[i-1][j-1] + 1

	m, n := len(nums1), len(nums2)

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化，这里如果是以(i,j)为下标，则需要对每行，没列单独初始化
	// 初始化列
	for i := 0; i < m; i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
	}

	// 初始化行
	for j := 0; j < n; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
		}
	}

	ans := 0

	// 状态转移方程
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] && i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}
