/// 零钱兑换
pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
    // 最大的硬币数是amout，因为硬币的面值最小是1
    let max_val = amount + 1;

    let amount = amount as usize;
    let mut dp = vec![max_val; amount + 1];
    dp[0] = 0;

    // 穷举每个目标数的最小硬币数
    for i in 1..=amount {
        //对每种子问题都穷举
        for v in coins.iter() {
            // v > i 的时候已经无法满足
            if *v as usize <= i {
                dp[i] = dp[i].min(dp[i - *v as usize] + 1);
            }
        }
    }

    if dp[amount] == max_val {
        return -1;
    }
    dp[amount]
}
