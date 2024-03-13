/// 括号生成
pub fn generate_parenthesis(n: i32) -> Vec<String> {
    // 满足条件的所有结果的合集
    let mut res = vec![];
    // 当前搜索过程中满足条件的路径
    let mut path = vec![];

    let n = n as usize;

    // symbol 当前左括号的个数，symbol < n , push (; if push ), symbol -1
    fn backtrack(symbol: usize, idx: usize, n: usize, path: &mut Vec<&str>, res: &mut Vec<String>) {
        if idx == n * 2 {
            if symbol == 0 {
                res.push(path.join(""));
            }
            return;
        }

        // symbol<n， 则可以push(
        if symbol < n {
            path.push("(");
            backtrack(symbol + 1, idx + 1, n, path, res);
            // backtrack
            path.pop();
        }

        if symbol > 0 {
            path.push(")");
            backtrack(symbol - 1, idx + 1, n, path, res);
            path.pop();
        }
    }

    backtrack(0, 0, n, &mut path, &mut res);

    res
}
