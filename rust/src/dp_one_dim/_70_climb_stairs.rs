/// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
/// 爬楼梯
/// dp
pub fn climb_stairs(n: i32) -> i32 {
    let mut dp = vec![0; n as usize + 1];
    dp[0] = 1;
    dp[1] = 1;

    for i in 2..=n as usize {
        dp[i] = dp[i - 1] + dp[i - 2];
    }

    return dp[n as usize];
}
