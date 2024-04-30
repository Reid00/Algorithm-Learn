///  最长连续递增序列
/// 有别于第300题最长递增序列，此题要连续递增
pub fn find_length_of_lcis(nums: Vec<i32>) -> i32 {
    // dp[i] 表示以i 为结尾的连续递增序列的长度是dp[i]
    // if dp[i] > dp[i-1], dp[i] = dp[i-1] + 1

    let n = nums.len();

    // 定义dp 并初始化
    let mut dp = vec![1; n];

    let mut ans = 1;

    for i in 1..n {
        if nums[i] > nums[i - 1] {
            dp[i] = dp[i - 1] + 1;
            ans = ans.max(dp[i]);
        }
    }

    ans
}

/// 贪心: 局部最优就是全局最优
pub fn find_length_of_lcis_(nums: Vec<i32>) -> i32 {
    let (mut win, mut ans) = (1, 1);

    for i in 1..nums.len() {
        if nums[i] > nums[i - 1] {
            win += 1;
            ans = ans.max(win);
        } else {
            win = 1;
        }
    }
    ans
}
