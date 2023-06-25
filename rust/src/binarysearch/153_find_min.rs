

pub fn find_min(nums: Vec<i32>) -> i32 {

    let (mut l, mut r) = (0, (nums.len()-1) as i32);

    while l <= r {

        let mid = l + ((r-l)>>1);

        if nums[mid as usize] < nums[r as usize] {
            // mid 有可能是最小值，不能是r = mid -1 会错过最小值
            r = mid;
        }else{
            l = mid + 1
        }
    }
    nums[r as usize]
}