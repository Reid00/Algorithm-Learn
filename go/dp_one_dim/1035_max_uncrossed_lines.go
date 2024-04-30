package dponedim

// maxUncrossedLines 不相交的线
// 经过分析和1143 本质一样，只要公共部分相对顺序相同则不会相交。变成了求最长公共子序列
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	// dp[i][j] = dp[i-1][j-1] + 1 if nums1[i-1]==nums2[j-1]
	// else dp[i][j] = max(dp[i-1][j], dp[i][j-1])

	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化略

	// 状态转移
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// 最长公共子序列一定出现在i,j 都是最大的时候
	return dp[m][n]
}
