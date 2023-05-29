pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {

    // 第一行，第一列 来标识某行某列出现0 
    let (n, m) = (matrix.len(), matrix[0].len());

    let (mut col0, mut row0) = (false, false);

    // 先判断第一行 和 第一列中是否有0
    // 遍历第一行元素，设置列标记
    for &i in matrix[0].iter() {
        if i == 0 {
            row0 = true;
        }
    }

    for i in matrix.iter() {
        if i[0] == 0 {
            col0 = true;
        }
    }


    // 遍历除去第一行 第一列的元素 如果有0，将对应的第一行第一列置为0
    for i in 1..n {
        for j in 1..m {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0;
                matrix[0][j] = 0;
            }
        }
    }

    // 根据重置的结果，遍历第一行 和 第一列数据
    // 第一列
    for i in 1..n {
        if matrix[i][0] == 0 {
            for j in 1..m {
                matrix[i][j] = 0;
            }
        }
    }

    for i in 1..m {
        if matrix[0][i] == 0 {
            for j in 1..n {
                matrix[j][i] = 0;
            }
        }
    }

    if row0 {
        for i in 0..m {
            matrix[0][i] = 0;
        }
    }

    if col0 {
        for i in 0..n {
            matrix[i][0] = 0;
        }
    }

}