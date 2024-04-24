/// 买卖股票的最佳时机含手续费
/// dp: 无交易次数限定
/// Time: O(n), Space: O(n)
pub fn max_profit(prices: Vec<i32>, fee: i32) -> i32 {
    // dp[i][j] 表示第i天，处于状态j 时的最大利润；此处i有两个值，0 持有股票和 1 不持有股票
    // dp[i][0]：第i天持有股票获得利润的最大值； dp[i][1]：第i天不持有股票获得利润的最大值；

    let n = prices.len();

    let mut dp = vec![vec![0; 2]; n];

    dp[0][0] = -prices[0];
    dp[0][1] = 0;

    for i in 1..n {
        dp[i][0] = dp[i - 1][0].max(dp[i - 1][1] - prices[i]);
        dp[i][1] = dp[i - 1][1].max(dp[i - 1][0] + prices[i] - fee);
    }

    dp[n - 1][1]
}
