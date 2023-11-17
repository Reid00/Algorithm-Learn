pub fn can_jump(nums: Vec<i32>) -> bool {
    let mut max_pos = 0;

    for (i, &v) in nums.iter().enumerate() {
        let i = i as i32;
        if i <= max_pos && i + v >= max_pos {
            max_pos = i + v;
        }
    }
    max_pos as usize >= nums.len() - 1
}