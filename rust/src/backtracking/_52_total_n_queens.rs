/// totalNQueens N 皇后II
/// 返回 n 皇后问题 不同的解决方案的数量。
pub fn total_n_queens(n: i32) -> i32 {
    let n = n as usize;
    let mut path = vec![vec!["."; n]; n];

    let mut res = 0;

    backtrack(&mut res, &mut path, 0, n);
    res
}

pub fn backtrack(res: &mut i32, path: &mut Vec<Vec<&str>>, row: usize, n: usize) {
    if row == n {
        *res += 1;
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

pub fn is_valid(path: &Vec<Vec<&str>>, r: usize, c: usize, n: usize) -> bool {
    for i in 0..n {
        for j in 0..n {
            if path[i][j] == "Q" && (j == c || i - j == r - c || i + j == r + c) {
                return false;
            }
        }
    }
    true
}
