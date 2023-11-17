fn find_middle_index(nums: Vec<i32>) -> i32 {
    let mut sum: i32 = nums.iter().sum();

    let mut left_sum = 0;

    for (i, &v) in nums.iter().enumerate() {
        if left_sum == sum - v {
            return i as i32;
        }

        left_sum += v;
        sum -= v;
    }

    return -1;
}
