/// 汇总区间
pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
    let n = nums.len();

    let mut result = vec![];

    let mut i = 0;

    while i < n {
        let left = i;

        i = i + 1;
        // 区间内的元素 相邻相差是1
        while i < n && nums[i - 1] + 1 == nums[i] {
            i += 1;
        }

        let res = if left < i - 1 {
            format!("{}->{}", nums[left], nums[i - 1])
        } else {
            nums[left].to_string()
        };

        result.push(res);
    }

    result
}
