/// N 皇后
pub fn solve_n_queens(n: i32) -> Vec<Vec<String>> {
    let n = n as usize;
    // 初始化棋盘, 同时作为搜索中满足条件的path
    let mut path = vec![vec!["."; n]; n];

    // 搜索结果
    let mut res = vec![];

    fn backtrack(res: &mut Vec<Vec<String>>, path: &mut Vec<Vec<&str>>, row: usize, n: usize) {
        // return condition
        if row == n {
            let mut p = vec![];
            for item in path {
                p.push(item.join(""));
            }
            res.push(p);
            return;
        }

        for col in 0..n {
            if is_valid(path, row, col, n) {
                path[row][col] = "Q";
                backtrack(res, path, row + 1, n);
                path[row][col] = ".";
            }
        }
    }

    backtrack(&mut res, &mut path, 0, n);
    res

}

/// 判断path[row][col] 是否可以放置Q
/// 同行、同列、对接线不可放置
fn is_valid(path: &Vec<Vec<&str>>, row: usize, col: usize, n: usize) -> bool {
    for i in 0..n {
        for j in 0..n {
            if path[i][j] == "Q" && (j == col || i + j == row + col || i - j == row - col) {
                return false;
            }
        }
    }
    true
}
