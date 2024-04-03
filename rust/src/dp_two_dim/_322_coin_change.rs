/// 零钱兑换
/// 返回能够兑换的最小硬币个数
/// 完全背包的视角再写一遍
pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
    // dp[j] 表示容量为j的背包填满，最少需要dp[j]个硬币
    // dp[j] = min(dp[j], dp[j-coins[i]] + 1)

    // 最大的硬币数是amout，因为硬币的面值最小是1
    let max_val = amount + 1;

    let amount = amount as usize;
    let mut dp = vec![max_val; amount + 1];
    dp[0] = 0;

    // 遍历物品
    for i in 0..coins.len() {
        // 遍历背包容量
        for j in 0..=amount {
            if j as i32 - coins[i] >= 0 {
                dp[j] = dp[j].min(dp[j - coins[i] as usize] + 1);
            }
        }
    }

    if dp[amount] == max_val {
        return -1;
    }
    dp[amount]
}
