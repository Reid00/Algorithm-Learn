/// combinationSum4 组合总和 Ⅳ
/// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
/// nums 中的元素可以重复使用
pub fn combination_sum4(nums: Vec<i32>, target: i32) -> i32 {
    let target = target as usize;

    let mut dp = vec![0; target + 1];
    dp[0] = 1;

    // 排列问题，外层是背包容量
    for i in 0..=target {
        for j in 0..nums.len() {
            if i as i32 - nums[j] >= 0 {
                dp[i] += dp[i - nums[j] as usize];
            }
        }
    }

    dp[target]
}
