



pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {

    // 二分
    let (mut l, mut r) = (0, (nums.len()-1) as i32);

    // 不用 <= 确保有两个元素
    while l < r {

        let mid = l + ((r-l)>>1);

        // 解释: 如果偶数和1 异或相当于+1， 奇数与1 异或相当于-1
        //  如果相邻的元素相同，说明mid 在单一元素x 的左侧
        if nums[mid as usize] == nums[(mid^1) as usize] {
            l = mid + 1;
        }else {
            r = mid;
        }
    }
    nums[l as usize]
}

// O(n)
pub fn single_non_duplicate2(nums: Vec<i32>) -> i32 {

    // 异或
    let mut ans = 0;

    for v in nums {
        ans ^= v;
    }
    ans
}