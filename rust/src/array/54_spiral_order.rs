
/// 螺旋矩阵
pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
    // 模拟
    let (rows, cols) = (matrix.len(), matrix[0].len());

    let mut visitd = vec![vec![false; cols]; rows];
    // 访问顺序的结果
    let total = rows * cols;
    let mut order = vec![0; total];
    // 右，下，左，上的顺序
    let directions: Vec<Vec<i32>> = vec![vec![0, 1], vec![1, 0], vec![0, -1], vec![-1, 0]];

    let mut direction_idx = 0;
    let (mut row, mut col) = (0, 0);

    for i in 0..total {
        order[i] = matrix[row][col];
        visitd[row][col] = true;

        let next_row = row as i32 + directions[direction_idx][0];
        let next_col = col as i32 + directions[direction_idx][1];
        // next_row, next_col 是否合法
        if next_row < 0
            || next_row >= rows as i32
            || next_col < 0
            || next_col >= cols as i32
            || visitd[next_row as usize][next_col as usize]
        {
            direction_idx = (direction_idx + 1) % 4;
        }
        // 经过上面的判断之后 row col as usize 是合法的
        row = (row as i32 + directions[direction_idx][0]) as usize;
        col = (col as i32 + directions[direction_idx][1]) as usize;
    }
    order
}
