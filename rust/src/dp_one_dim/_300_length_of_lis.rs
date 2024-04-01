/// 最长递增子序列
pub fn length_of_lis(nums: Vec<i32>) -> i32 {
    let n = nums.len();

    // 初始的递增子序列长度都是1
    // dp[i] 表示从开始到i结尾的最长递增子序列长是为dp[i], 包含i
    // dp[i] = max(dp[i], dp[j] + 1) 0<=j <i, nums[j] < nums[i]
    let mut dp = vec![1; n];

    for i in 0..n {
        for j in 0..i {
            if nums[j] < nums[i] {
                dp[i] = dp[i].max(dp[j] + 1);
            }
        }
    }

    *dp.iter().max().unwrap()
}
