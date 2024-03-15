/// 环形子数组的最大和
pub fn max_subarray_sum_circular(nums: Vec<i32>) -> i32 {
    // kadane 算法
    // 最大子数组 成环 --> 看下图，总和减去不成环的子数组最小和(total - minSum) === 成环区域的最大和
    let mut min_ending_here = nums[0];
    let mut max_ending_here = nums[0];

    let mut min_so_far = nums[0];
    let mut max_so_far = nums[0];

    for i in 1..nums.len() {
        max_ending_here = (max_ending_here + nums[i]).max(nums[i]);
        min_ending_here = (min_ending_here + nums[i]).min(nums[i]);

        max_so_far = max_so_far.max(max_ending_here);
        min_so_far = min_so_far.min(min_ending_here);
    }

    let total: i32 = nums.iter().sum();

    if total - min_so_far > 0 {
        return max_so_far.max(total - min_so_far);
    }
    max_so_far
}
