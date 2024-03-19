/// 搜索二维矩阵
pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
    let (m, n) = (matrix.len(), matrix[0].len());

    let (mut l, mut r) = (0, m * n - 1);

    while l <= r {
        let mid = l + ((r - l) >> 1);

        let (row, col) = (mid / n, mid % n);
        if matrix[row][col] == target {
            return true;
        } else if matrix[row][col] < target {
            l = mid + 1
        } else {
            // mid 是usize， 防止mid=0 的时候 mid -1 越界
            if mid == 0 {
                return false;
            }
            r = mid - 1
        }
    }

    false
}
