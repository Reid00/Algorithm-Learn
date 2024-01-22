/// 生命游戏
pub fn game_of_life(board: &mut Vec<Vec<i32>>) {
    // METHOD 1 copy 一份board，遍历元素时对每个元素的8个邻居进行判断然后修改copyed board value
    // METHOD 2: O(1) 空间复杂度
    // 引入复合状态 -1, 表示 原先是1， 但是规则1、3 导致变成0，先用1 替代
    // 2, 表示 原先是0，规则4 导致变成了1， 先用2 替代

    let (m, n) = (board.len(), board[0].len());

    let directions = vec![
        vec![0, 1],
        vec![0, -1],
        vec![-1, 0],
        vec![1, 0],
        vec![1, 1],
        vec![1, -1],
        vec![-1, 1],
        vec![-1, -1],
    ];

    for i in 0..m {
        for j in 0..n {
            // 元素周围8个方向的 1 的个数
            let mut nums = 0;
            for direction in directions.iter() {
                let (next_row, next_col) = (i as i32 + direction[0], j as i32 + direction[1]);

                if next_row >= 0
                    && next_row < m as i32
                    && next_col >= 0
                    && next_col < n as i32
                    && board[next_row as usize][next_col as usize].abs() == 1
                {
                    nums += 1;
                }
            }

            match board[i][j] {
                1 => {
                    // 规则1,3
                    if nums < 2 || nums > 3 {
                        board[i][j] = -1;
                    }
                }
                _ => {
                    // 规则4
                    if nums == 3 {
                        board[i][j] = 2;
                    }
                }
            }
        }
    }

    for rows in board.iter_mut() {
        for v in rows.iter_mut() {
            if *v > 0 {
                *v = 1;
            } else {
                *v = 0;
            }
        }
    }
}
