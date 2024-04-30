/// 不相交的线
/// 经过分析和1143 本质一样，只要公共部分相对顺序相同则不会相交。变成了求最长公共子序列
pub fn max_uncrossed_lines(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
    // dp[i][j] = dp[i-1][j-1] + 1 if nums1[i-1]==nums2[j-1]
    // else dp[i][j] = max(dp[i-1][j], dp[i][j-1])

    let (m, n) = (nums1.len(), nums2.len());

    let mut dp = vec![vec![0; n + 1]; m + 1];

    // 初始化，略

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if nums1[i - 1] == nums2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = dp[i - 1][j].max(dp[i][j - 1]);
            }
        }
    }

    dp[m][n]
}
