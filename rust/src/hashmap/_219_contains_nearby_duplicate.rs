///  存在重复元素 II
pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
    // HashMap k: nums 中的元素， v: 该元素中最大的index
    // 为什么这个v: 因为要两个元素距离<=k 只需要记录最大的idx，待后续出现相同的元素和最大的idx 比较即可
    use std::collections::HashMap;

    let mut hmap = HashMap::new();

    let found = nums.iter().enumerate().any(|(idx, &v)| {
        if let Some(old_idx) = hmap.insert(v, idx) {
            if idx - old_idx <= k as usize {
                return true;
            }
        }
        false
    });
    found
}

pub fn contains_nearby_duplicate2(nums: Vec<i32>, k: i32) -> bool {
    // 滑动窗口，长度为k 的窗口内 如果存在两个相同的元素则返回true

    use std::collections::HashSet;
    let mut set = HashSet::new();

    let found = nums.iter().enumerate().any(|(i, &v)| {
        // i 超过k 的时候 每次右移都要清除窗口左侧的元素
        if i > k as usize {
            set.remove(&nums[i - k as usize - 1]);
        }
        if !set.insert(v) {
            return true;
        }
        false
    });
    found
}
