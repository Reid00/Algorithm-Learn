
/// 乘积最大子数组
pub fn max_product(nums: Vec<i32>) -> i32 {
    let (mut max_ending_here, mut min_ending_here) = (nums[0], nums[0]);
    let mut max_so_far = nums[0];

    for i in 1..nums.len() {
        let old_max = max_ending_here;

        max_ending_here = nums[i]
            .max(max_ending_here * nums[i])
            .max(min_ending_here * nums[i]);

        min_ending_here = nums[i]
            .min(min_ending_here * nums[i])
            .min(old_max * nums[i]);

        max_so_far = max_so_far.max(max_ending_here);
    }

    max_so_far
}
