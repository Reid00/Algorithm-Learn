/// 买卖股票的最佳时机
/// dp: 你最多可以完成 两笔 交易
/// Time: O(n), Space: O(n)
pub fn max_profit(prices: Vec<i32>) -> i32 {
    // dp[i][j] 表示第i天，处于j状态时的最大利润,
    // 因为可以买卖两次，j 有四种状态
    // j=1 第一次持有
    // j=2 第一次不持有
    // j=3 第二次持有
    // j=4 第二次不持有
    // dp[i][1] 表示第i天，第一次持有获取的最大利润, 如果i-1 已经持有,dp[i][1] = dp[i-1][1]
    // 如果i-1 没有持有，dp[i][1] = -prices[i]
    // dp[i][2] 表示第i天，第一次不持有获取的最大利润，如果i-1 已经持有，需要卖出,dp[i][2] = dp[i-1][1] + prices[i]
    // 如果i-1 没有持有，dp[i][2] = dp[i-1][2]
    // dp[i][3] 表示第i天，第二次持有获取的最大利润, 如果i-1 已经持有，dp[i][3] = dp[i-1][3]
    // 如果i-1 没有持有，dp[i][3] = dp[i-1][2] - prices[i]
    // dp[i][4] 表示第i天，第二次不持有获取的最大利润，if i-1 已经持有，dp[i][4] = dp[i-1][3] + prices[i]
    // 如果i-1 没有持有, dp[i][4]= dp[i-1][4]

    let n = prices.len();

    let mut dp = vec![vec![0; 5]; n];

    // 初始化dp
    dp[0][1] = -prices[0];
    dp[0][2] = 0;
    dp[0][3] = -prices[0];
    dp[0][4] = 0;

    for i in 1..n {
        dp[i][1] = dp[i - 1][1].max(-prices[i]);
        dp[i][2] = dp[i - 1][2].max(dp[i - 1][1] + prices[i]);
        dp[i][3] = dp[i - 1][3].max(dp[i - 1][2] - prices[i]);
        dp[i][4] = dp[i - 1][4].max(dp[i - 1][3] + prices[i]);
    }

    dp[n - 1][4]
}
