/// 买卖股票的最佳时机
/// 贪心算法: 不限买卖次数
/// Time: O(n), Space: O(1)
pub fn max_profit(prices: Vec<i32>) -> i32 {
    // dp[i][j] 表示第i天 持有(j=0) 和不持有(j=1) 股票时的最大利润
    // dp[i][0] 表示第i天，持有股票时的最大利润; 如果第i-1天持有股票，则i天维持即可,dp[i][0] = dp[i-1][0]
    // 如果第i-1天 不持有股票，因为可以多次买卖, 第i天的利润为第i-1天持有的减i天的消耗, dp[i][0] = dp[i-1][1] - prices[i]
    // dp[i][1] 表示第i天，不持有股票的最大利润; 如果第i-1天持有股票，则i天卖出,dp[i][1] = dp[i-1][0] + prices[i]
    // 如果第i-1天 不持有股票，则维持，dp[i][1] = dp[i-1][1]

    let n = prices.len();

    let mut dp = vec![vec![0; 2]; n];

    // 初始化
    dp[0][0] = -prices[0]; //第0天持有股票的利润
    dp[0][1] = 0; // 第0天 不持有股票的利润

    for i in 1..n {
        dp[i][0] = dp[i - 1][0].max(dp[i - 1][1] - prices[i]);
        dp[i][1] = dp[i - 1][1].max(dp[i - 1][0] + prices[i]);
    }

    dp[n - 1][1]
}

/// 买卖股票的最佳时机
/// 贪心算法: 不限买卖次数
/// Time: O(n), Space: O(1)
pub fn max_profit2(prices: Vec<i32>) -> i32 {
    prices
        .windows(2)
        .map(|x| x[1] - x[0])
        .filter(|x| *x > 0)
        .sum()
}
