/// 买卖股票的最佳时机
/// dp: 你最多可以完成 k笔 交易
/// Time: O(n), Space: O(n)
pub fn max_profit(k: i32, prices: Vec<i32>) -> i32 {
    let n = prices.len();

    let mut dp = vec![vec![0; (2 * k + 1) as usize]; n];

    // 初始化第0天最大利润; 奇数持有，偶数不持有
    for j in 1..(2 * k + 1) as usize {
        if j & 1 == 1 {
            dp[0][j] = -prices[0];
        } else {
            dp[0][j] = 0;
        }
    }

    for i in 1..n {
        for j in 1..(2 * k + 1) as usize {
            // 奇数持有股票时的最大利润
            if j & 1 == 1 {
                dp[i][j] = dp[i - 1][j].max(dp[i - 1][j - 1] - prices[i]);
            } else {
                dp[i][j] = dp[i - 1][j].max(dp[i - 1][j - 1] + prices[i]);
            }
        }
    }

    dp[n - 1][2 * k as usize]
}
