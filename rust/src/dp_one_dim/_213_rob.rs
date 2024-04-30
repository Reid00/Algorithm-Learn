/// 打家劫舍 II
/// 第一个房屋和最后一个房屋是紧挨着的
pub fn rob(nums: Vec<i32>) -> i32 {
    // 因为第一个房间和最后一个房间首位相连 so 如果偷窃了第一间房屋，则不能偷窃最后一间房屋
    // 因此分两种范围偷第一个: [0, n-2] 和 不偷第一个：[1, n-1]

    let n = nums.len();

    if n == 1 {
        return nums[0];
    }

    if n == 2 {
        return nums[0].max(nums[1]);
    }

    let res1 = rob_range(&nums, 0, n - 2);
    let res2 = rob_range(&nums, 1, n - 1);

    res1.max(res2)
}

pub fn rob_range(nums: &[i32], start: usize, end: usize) -> i32 {
    let mut dp = vec![0; nums.len()];

    dp[start] = nums[start];
    dp[start + 1] = nums[start].max(nums[start + 1]);

    for i in start + 2..=end {
        dp[i] = dp[i - 1].max(dp[i - 2] + nums[i]);
    }

    dp[end]
}
