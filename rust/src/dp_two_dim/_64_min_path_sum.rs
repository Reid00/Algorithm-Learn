/// 最小路径和
/// 每次只能向下或者向右移动一步
pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
    // dp[i][j] 表示走到格子i,j 位置的，最小路径和
    // 到i,j 这个位置，可以从(i,j-1) 或者(i-1,j) 移动而来
    // dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]

    let (n, m) = (grid.len(), grid[0].len());

    let mut dp = vec![vec![0; m]; n];

    // init dp
    dp[0][0] = grid[0][0];

    // 最左侧的列只能从上面走过来
    for i in 1..n {
        dp[i][0] = dp[i - 1][0] + grid[i][0];
    }

    // 最上面的行，只能从左侧走过来
    for j in 1..m {
        dp[0][j] = dp[0][j - 1] + grid[0][j];
    }

    for i in 1..n {
        for j in 1..m {
            dp[i][j] = dp[i - 1][j].min(dp[i][j - 1]) + grid[i][j];
        }
    }

    dp[n - 1][m - 1]
}
