///最大字数组和
pub fn max_sub_array(nums: Vec<i32>) -> i32 {
    let n = nums.len();
    let mut dp = vec![0; n];
    dp[0] = nums[0];

    for i in 1..n {
        if dp[i - 1] < 0 {
            dp[i] = nums[i]
        } else {
            dp[i] = dp[i - 1] + nums[i]
        }
    }

    dp.into_iter().max().unwrap()
}
