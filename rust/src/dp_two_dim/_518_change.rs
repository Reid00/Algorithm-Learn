/// change 零钱兑换II
/// 计算并返回可以凑成总金额的硬币组合数
/// 完全背包视角 重新解读此题
pub fn change(amount: i32, coins: Vec<i32>) -> i32 {
    // dp[j] 表示填满容量为j 的背包，有dp[j] 中方法
    // dp[j] += dp[j-coins[i]]

    let amount = amount as usize;
    // 定义dp
    let mut dp = vec![0; amount + 1];
    // 初始化dp
    dp[0] = 1;

    for i in 0..coins.len() {
        for j in coins[i] as usize..=amount {
            dp[j] += dp[j - coins[i] as usize];
        }
    }

    dp[amount]
}
