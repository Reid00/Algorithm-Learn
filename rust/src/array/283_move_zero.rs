


pub fn move_zeroes(nums: &mut Vec<i32>) {
    let (mut left, mut right) = (0, 0);

    while right < nums.len() {
        if nums[right] != 0 {
           nums.swap(left, right);
           left += 1;
        }
        right +=1;
    }

}