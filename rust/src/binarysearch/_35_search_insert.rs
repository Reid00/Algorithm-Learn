/// 搜索插入位置
pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    let (mut l, mut r) = (0, nums.len());

    while l <= r {
        let mid = l + ((r - l) >> 1);

        if nums[mid] == target {
            return mid as i32;
        } else if nums[mid] > target {
            // mid usize 为0, mid - 1 会溢出
            if mid == 0 {
                return mid as i32;
            }
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }

    l as i32
}

/// 搜索插入位置
pub fn search_insert2(nums: Vec<i32>, target: i32) -> i32 {
    nums.binary_search(&target).unwrap_or_else(|x| x) as i32
}
