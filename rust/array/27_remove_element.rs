


fn remove_elements(nums: &mut Vec<i32>, val: i32) -> i32 {
    
    // 快慢指针
    let mut slow = 0;

    for fast in 0..nums.len() {

        if nums[fast] != val {
            nums[slow] = nums[fast];
            slow += 1;
        }
    }

    slow as i32

    // 自带API
//     nums.retain(|&x| x != val);
//     nums.len() as i32
}