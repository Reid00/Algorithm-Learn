/// change 零钱兑换II
/// 计算并返回可以凑成总金额的硬币组合数
pub fn change(amount: i32, coins: Vec<i32>) -> i32 {
    let amount = amount as usize;

    let mut dp = vec![0; amount + 1];
    dp[0] = 1;

    for &coin in coins.iter() {
        for i in 1..=amount {
            if i >= coin as usize {
                dp[i] += dp[i - coin as usize]
            }
        }
    }
    dp[amount]
}
