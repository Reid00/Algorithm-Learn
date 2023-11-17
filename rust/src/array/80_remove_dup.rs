pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
    let length = nums.len();

    if length <= 2 {
        return length as i32;
    }

    let mut slow = 2;

    for fast in 2..nums.len() {
        if nums[fast] != nums[slow - 2] {
            nums[slow] = nums[fast];
            slow += 1;
        }
    }
    slow as i32
}
