/// 插入区间
pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
    // 1. 申请新的数组，其中有个位置i 是要插入的位置，预先假设放入的是 new_interval 元素
    // 2. if v[1] < new_interval[0]， v 在i 的左侧，直接push
    // 3. if v[0] > new_interval[1], v 在 i 的右侧，先保证 i 已经插入元素
    // 4. 其他情况表明 和要插入的元素new_interval 和交集，先进行merge

    let mut res = vec![];
    let (mut left, mut right) = (new_interval[0], new_interval[1]);
    let mut merged = false;

    for v in intervals {
        // 在 i 的左侧
        if v[1] < left {
            res.push(v);
        } else if v[0] > right {
            // 在 i 的右侧
            // 先保证i位置已经插入了
            if !merged {
                res.push(vec![left, right]);
                merged = true;
            }

            res.push(v);
        } else {
            // 合并
            left = left.min(v[0]);
            right = right.max(v[1]);
        }
    }

    if !merged {
        res.push(vec![left, right]);
    }

    res
}
