// 最后一块石头的重量 II
pub fn last_stone_weight_ii(stones: Vec<i32>) -> i32 {
    // dp[j] 表示容量为j 的背包，最大的容纳重量是dp[j]

    let sum: i32 = stones.iter().sum();

    let bag_size = sum / 2;

    let mut dp = vec![0; bag_size as usize + 1];

    for i in 0..stones.len() {
        for j in (stones[i]..=bag_size).rev() {
            if j >= stones[i] {
                dp[j as usize] = dp[j as usize].max(dp[(j - stones[i]) as usize] + stones[i]);
            }
        }
    }

    sum - dp[bag_size as usize] * 2
}
