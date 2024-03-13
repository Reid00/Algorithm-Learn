/// combine 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
    // 所有满足条件的集合
    let mut res = vec![];
    // 当前满足条件的路径
    let mut path = vec![];

    fn backtrack(start_idx: i32, n: i32, k: i32, path: &mut Vec<i32>, res: &mut Vec<Vec<i32>>) {
        if path.len() == k as usize {
            res.push(path.clone());
            return;
        }

        for i in start_idx..=(n - (k - path.len() as i32) + 1) {
            path.push(i);
            backtrack(i + 1, n, k, path, res);
            path.pop();
        }
    }

    backtrack(1, n, k, &mut path, &mut res);
    res
}
