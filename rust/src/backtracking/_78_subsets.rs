/// 子集
pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut res = vec![];
    let mut path = vec![];

    fn backtrack(start_idx: usize, nums: &[i32], res: &mut Vec<Vec<i32>>, path: &mut Vec<i32>) {
        res.push(path.clone());

        for i in start_idx..nums.len() {
            path.push(nums[i]);
            backtrack(i + 1, nums, res, path);
            path.pop();
        }
    }

    backtrack(0, &nums, &mut res, &mut path);
    res
}
