/// 最长重复子数组(连续的)
pub fn find_length(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
    // dp[i][j] 表示以nums1[i-1] 下标 和 以nums2[j-1] 为下标的最长重复字数组的长度是dp[i][j]
    // 注意此处是（i-1,j-1）为下标，不是(i,j) 为下标，是为了方便初始化
    // if nums1[i-1] == nums2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
    // 因为dp[i][j] 的意义是下标(i-1, j-1) 所以dp长度是 m + 1 * n+1

    let (m, n) = (nums1.len(), nums2.len());

    let mut dp = vec![vec![0; n + 1]; m + 1];

    // 初始化
    // dp[0][0] 没有太大意义，都为0

    let mut ans = 0;

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if nums1[i - 1] == nums2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1;
                ans = ans.max(dp[i][j]);
            }
        }
    }

    ans
}
