/// 买卖股票的最佳时机
/// dp: 股票就买卖一次
/// Time: O(n), Space: O(n)
pub fn max_profit(prices: Vec<i32>) -> i32 {
    // dp[i][j]	表示第i天，现在有两种状态, j=0,表示持有股票; j=1，表示不持有股票
    // so dp[i][0] 表示第i天，持有股票能获得的最大利润
    // dp[i][1] 表示第i天，不持有股票能获取的最大利润
    // 递推公式
    // dp[i][0] 第i天 持有股票，有两种情况，第i-1天，如果持有股票，因为只能一次买卖，所以dp[i][0] = dp[i-1][0]
    // 如果第i-1 天，没有持有股票，dp[i][0] = - prices[i]
    // 即: dp[i][0] = max(dp[i-1][0], -prices[i])
    // dp[i][1] 第i天，不持有股票，有两种情况，第i-1天，本身就不持有股票，则dp[i][1]= dp[i-1][1]
    // 如果第i-1 天，持有股票，那第i天就需要进行卖掉，dp[i][1] = dp[i-1][0] + prices[i]
    // 即: dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i])

    let n = prices.len();
    let mut dp = vec![vec![0; 2]; n];

    dp[0][0] = -prices[0];
    dp[0][1] = 0;

    for i in 1..n {
        dp[i][0] = dp[i - 1][0].max(-prices[i]);
        dp[i][1] = dp[i - 1][1].max(dp[i - 1][0] + prices[i]);
    }

    dp[n - 1][1]
}

/// 买卖股票的最佳时机
/// 贪心算法: 因为股票就买卖一次，可以寻找最低点和最高点进行买卖
/// Time: O(n), Space: O(1)
pub fn max_profit2(prices: Vec<i32>) -> i32 {
    let mut res = 0;
    let mut low = i32::MAX;

    for v in prices {
        low = low.min(v);
        res = res.max(v - low);
    }
    res
}
