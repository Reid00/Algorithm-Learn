/// 单词搜索
pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
    if board.is_empty() {
        return false;
    }

    let row = board.len();
    let col = board[0].len();

    // visited 当前搜索时是否被访问过
    let mut visited = vec![vec![false; col]; row];
    // word 转化为 Vec
    let word: Vec<char> = word.chars().collect();

    let dirs = vec![vec![0, 1], vec![0, -1], vec![-1, 0], vec![1, 0]];

    for r in 0..row {
        for c in 0..col {
            if backtrack(&board, r, c, 0, &dirs, &word, &mut visited) {
                return true;
            }
        }
    }

    false
}

fn backtrack(
    board: &Vec<Vec<char>>,
    r: usize,
    c: usize,
    idx: usize,
    dirs: &Vec<Vec<i32>>,
    word: &Vec<char>,
    visited: &mut Vec<Vec<bool>>,
) -> bool {
    if board[r][c] != word[idx] {
        return false;
    }

    if idx == word.len() - 1 {
        return true;
    }

    if board[r][c] == word[idx] {
        visited[r][c] = true;

        for dir in dirs {
            let (x, y) = ((r as i32 + dir[0]), (c as i32 + dir[1]));
            if in_area(board, x, y) && !visited[x as usize][y as usize] {
                if backtrack(board, x as usize, y as usize, idx + 1, dirs, word, visited) {
                    return true;
                }
            }
        }
        // backtrack
        visited[r][c] = false;
    }

    false
}

/// 判断坐标x,y 是否在board 内
fn in_area(board: &Vec<Vec<char>>, x: i32, y: i32) -> bool {
    x >= 0 && x < board.len() as i32 && y >= 0 && y < board[0].len() as i32
}
