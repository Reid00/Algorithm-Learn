/// 组合总和
pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    let mut res = vec![];
    let mut path = vec![];

    let mut candidates = candidates;

    fn backtrack(
        start_idx: usize,
        mut total: i32,
        target: i32,
        path: &mut Vec<i32>,
        res: &mut Vec<Vec<i32>>,
        candidates: &[i32],
    ) {
        if total > target {
            return;
        }

        if total == target {
            res.push(path.clone());
            return;
        }

        for i in start_idx..candidates.len() {
            // 如果candidates 是排好序的，剪枝
            if total > target {
                break;
            }

            path.push(candidates[i]);
            backtrack(i, total + candidates[i], target, path, res, candidates);
            // 回溯
            path.pop();
        }
    }

    candidates.sort();
    backtrack(0, 0, target, &mut path, &mut res, &candidates);
    res
}
