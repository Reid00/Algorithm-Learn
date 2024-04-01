/// 斐波那契数
pub fn fib(n: i32) -> i32 {
    if n < 2 {
        return n;
    }

    /// dp 优化
    let n = n as usize;

    let mut dp = vec![0; n + 1];
    dp[0] = 0;
    dp[1] = 1;

    for i in 2..=n {
        dp[i] = dp[i - 1] + dp[i - 2];
    }

    dp[n]
}
