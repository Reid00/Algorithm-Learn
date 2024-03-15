pub fn max_sub_array(nums: Vec<i32>) -> i32 {
    // Kadane 算法
    if nums.is_empty() {
        return 0;
    }

    let (mut max_ending_here, mut max_so_far) = (nums[0], nums[0]);

    for v in nums.into_iter().skip(1) {
        max_ending_here = (max_ending_here + v).max(v);
        max_so_far = max_so_far.max(max_ending_here);
    }

    max_so_far
}

pub fn max_sub_array2(nums: Vec<i32>) -> i32 {
    // 前缀和 => pre_sum, pre_sum[0] = 0
    // pre_sum[i+1] = pre_sum[i] + nums[i]
    // pre_sum[i] 表示 nums[0..i-1] 的子数组和
    // nums[i..j] 的字数组和 => pre_sum[j+1] - pre_sum[j]
    // 前缀和
    let mut pre_sum = 0;
    // 最小的前缀和
    let mut min_pre_sum = 0;
    let mut ans = i32::MIN;

    for v in nums {
        // 求出到v 为止的前缀和
        pre_sum += v;
        // 计算每次和最小前缀和的差值 => 前缀和的差值即是字数组的和
        ans = ans.max(pre_sum - min_pre_sum);
        // update
        min_pre_sum = min_pre_sum.min(pre_sum);
    }
    ans
}
