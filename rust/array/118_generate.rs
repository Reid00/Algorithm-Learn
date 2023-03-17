


// 杨辉三角
pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {

    let mut result: Vec<Vec<i32>> = vec![];

    for i in 0..num_rows as usize {

        // 初始把一行数据全部赋值为1
        result.push(vec![1;i+1]);

        for j in 1..i {
            // 根据杨辉三角赋值
            result[i][j] = result[i-1][j-1] + result[i-1][j]
        }
    }

    result

}