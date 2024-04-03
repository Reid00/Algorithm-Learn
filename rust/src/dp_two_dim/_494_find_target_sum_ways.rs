/// 目标和
pub fn find_target_sum_ways(nums: Vec<i32>, target: i32) -> i32 {
    let sum: i32 = nums.iter().sum();

    if target.abs() > sum {
        return 0;
    }

    // 不是偶数
    if (sum + target) & 1 == 1 {
        return 0;
    }

    let bag_size = (sum + target) >> 1;

    let mut dp = vec![0; bag_size as usize + 1];
    dp[0] = 1;

    for i in 0..nums.len() {
        // 先遍历物品
        for j in (nums[i]..=bag_size).rev() {
            // 遍历容量
            if j >= nums[i] {
                dp[j as usize] += dp[(j - nums[i]) as usize];
            }
        }
    }

    dp[bag_size as usize]
}
