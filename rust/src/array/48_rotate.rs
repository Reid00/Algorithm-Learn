


// 旋转图像(二维数组)
pub fn rotate(matrix: &mut Vec<Vec<i32>>) {

    // 水平翻转 + 对角线翻转
    let n = matrix.len();

    for i in 0..n/2 {
        // (&mut matrix[i], &mut matrix[n-i-1]) = (&mut matrix[n-i-1], &mut matrix[i]);
        matrix.swap(i, n-i-1);
    }

    for i in 0..n {
        for j in 0..i {
            (matrix[i][j], matrix[j][i]) = (matrix[j][i], matrix[i][j]);
            
        }
    }

}