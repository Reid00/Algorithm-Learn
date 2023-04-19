

pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {

    let mut rows = vec![vec![0;9];9];
    let mut cols = vec![vec![0;9];9];

    // 注意此处理解为 三个 3 * 9 的二维数组
    let mut subboxs = vec![vec![vec![0;9];3];3];

    for (i, row) in board.iter().enumerate() {
        for (j, &c) in row.iter().enumerate() {
            if c == '.' {
                continue;
            }

            let index = (c as usize - '1' as usize);

            rows[i][index] += 1;
            cols[j][index] += 1;
            subboxs[i/3][j/3][index] +=1;

            if rows[i][index] > 1 || cols[j][index] > 1 || subboxs[i/3][j/3][index] > 1 {
                return false
            }
        }
    }

    true
}