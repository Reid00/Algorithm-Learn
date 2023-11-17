pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    match nums.binary_search(&target) {
        Ok(x) => return x as i32,
        Err(x) => return x as i32,
    }
}

pub fn search_insert2(nums: Vec<i32>, target: i32) -> i32 {
    nums.binary_search(&target).unwrap_or_else(|x| x) as i32
}

pub fn search_insert3(nums: Vec<i32>, target: i32) -> i32 {
    let n = nums.len();

    let (mut l, mut r) = (0, n - 1);
    // 0, 3
    let mut ans = n;

    while l <= r {
        let mid = l + ((r - l) >> 1);

        if nums[mid] == target {
            return mid as i32;
        }

        if nums[mid] < target {
            l = mid + 1;
        } else {
            ans = mid;
            // 此处mid 是usize 类型，如果mid == 0 ，r = mid -1 会越界 需要注意
            if mid == 0 {
                return ans as i32;
            }
            r = mid - 1;
        }
    }

    return ans as i32;
}
