/// 全排列
pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut res = vec![];
    let mut path = vec![];

    fn backtrack(nums: &[i32], path: &mut Vec<i32>, res: &mut Vec<Vec<i32>>) {
        if path.len() == nums.len() {
            res.push(path.clone());
            return;
        }

        for item in nums.iter() {
            if !path.contains(item) {
                path.push(*item);
                backtrack(nums, path, res);
                path.pop();
            }
        }
    }

    backtrack(&nums, &mut path, &mut res);
    res
}
