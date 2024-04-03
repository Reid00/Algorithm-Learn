/// 完全平方数
/// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
/// 完全平方数 是一个整数，其值等于另一个整数的平方
/// 其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
pub fn num_squares(n: i32) -> i32 {
    // 完全背包的思维
    // dp[j] 表示要凑够j的值，需要完全平方数至少dp[j]个
    // dp[j] = min(dp[j], dp[j-i*i] + 1)

    let n = n as usize;
    let max_val = n + 1;
    // 初始化,dp[0]=0
    // 其他为max_val
    let mut dp = vec![max_val; n + 1];
    dp[0] = 0;

    // 遍历物品, 物品要 i * i <= n
    for i in 0..=(n as f64).sqrt() as usize {
        // 遍历背包
        for j in i * i..=n {
            dp[j] = dp[j].min(dp[j - i * i] + 1);
        }
    }

    if dp[n] == n + 1 {
        return -1;
    }
    dp[n] as i32
}
