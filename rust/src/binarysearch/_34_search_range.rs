// searchRange 在排序数组(非递减)中查找元素的第一个和最后一个位置
pub fn search_range2(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let n = nums.len() as i32;

    let (mut l, mut r) = (0, n - 1);

    let mut ans = vec![];

    while l <= r {
        let mid = l + ((r - l) >> 1);

        if nums[mid as usize] == target {
            // 向左找， 找到相等的起始位置

            let mut left = mid as i32;
            while left >= 0 && nums[left as usize] == target {
                left -= 1;
            }
            // 退出上面循环时，left 是满足相等的条件下-1 了
            ans.push(left + 1);

            let mut right = mid;

            while right < n && nums[right as usize] == target {
                right += 1;
            }

            ans.push(right - 1);
            return ans;
        }

        if nums[mid as usize] > target {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
    vec![-1, -1]
}
