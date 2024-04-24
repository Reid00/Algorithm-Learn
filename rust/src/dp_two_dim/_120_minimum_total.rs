/// 三角形最小路径和
pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
    // dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]

    let n = triangle.len();

    let mut dp = vec![vec![0; n]; n];

    // 初始化dp
    dp[0][0] = triangle[0][0];

    for i in 1..n {
        // 最左侧
        dp[i][0] = dp[i - 1][0] + triangle[i][0];

        for j in 1..i {
            dp[i][j] = dp[i - 1][j].min(dp[i - 1][j - 1]) + triangle[i][j];
        }

        // 最右侧
        dp[i][i] = dp[i - 1][i - 1] + triangle[i][i];
    }

    let mut res = i32::MAX;

    for i in 0..n {
        res = res.min(dp[n - 1][i]);
    }
    res
}
