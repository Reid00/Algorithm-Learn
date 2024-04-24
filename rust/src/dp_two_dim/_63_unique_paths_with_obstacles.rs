/// 不同路径 II
/// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
/// 网格中的障碍物和空位置分别用 1 和 0 来表示。
pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
    // dp[i][j]	表示到达(i,j) 位置总共有dp[i][j] 条路径
    // dp[i][j] = dp[i-1][j] + dp[i][j-1] 意思是 到(i,j) 的路径条数是 到(i-1,j) 和 (i,j-1)
    // 的路径之后 在到(i,j)

    let (n, m) = (obstacle_grid.len(), obstacle_grid[0].len());

    let mut dp = vec![vec![0; m]; n];

    dp[0][0] = 0;

    // 最左边的列
    for i in 1..n {
        if obstacle_grid[i][0] != 1 {
            dp[i][0] = 1;
        } else {
            dp[i][0] = 0;

            for x in i + 1..n {
                dp[x][0] = 0;
            }
            break;
        }
    }

    // 最上面的行
    for j in 1..m {
        if obstacle_grid[0][j] != 1 {
            dp[0][j] = 1;
        } else {
            dp[0][j] = 0;

            for x in j + 1..m {
                dp[0][x] = 0;
            }
            break;
        }
    }

    for i in 1..n {
        for j in 1..m {
            if obstacle_grid[i][j] != 1 {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            } else {
                dp[i][j] = 0;
            }
        }
    }

    dp[n - 1][m - 1]
}
