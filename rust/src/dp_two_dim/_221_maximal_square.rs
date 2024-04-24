/// 最大正方形
/// dp Time: O(m*n), Space: O(m*n)
pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
    // dp[i][j] 表示以(i,j)坐标且值包含 1 的正方形的最大边长
    // 状态转移方程:
    // 以(i,j) 为右下角正方形边长受限于其正上方，左上方，左边正方形的边长的最小值
    // dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

    let (m, n) = (matrix.len(), matrix[0].len());

    let mut dp = vec![vec![0; n]; m];
    // 记录最大的变长
    let mut max_side = 0;

    // 初始化
    for i in 0..m {
        for j in 0..n {
            dp[i][j] = matrix[i][j] as i32 - '0' as i32;
            max_side = max_side.max(dp[i][j]);
        }
    }

    for i in 1..m {
        for j in 1..n {
            if dp[i][j] == 1 {
                dp[i][j] = dp[i - 1][j - 1].min(dp[i - 1][j]).min(dp[i][j - 1]) + 1;
                max_side = max_side.max(dp[i][j]);
            }
        }
    }

    max_side * max_side
}
