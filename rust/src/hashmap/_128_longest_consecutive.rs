/// 最长连续序列
pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
    // 先用hashMap 去重，
    // 找到连续序列中的最小值v(v-1 不在hmap 里面则是最小值)
    // 循环 v+1/v+2/... 查看是否在hmap 中

    use std::collections::HashSet;

    let mut set = HashSet::new();

    nums.iter().for_each(|&v| {
        set.insert(v);
    });

    let mut result = 0;
    set.iter().for_each(|&x| {
        if !set.contains(&(x - 1)) {
            let mut cur_num = x;
            let mut length = 1;

            while set.contains(&(cur_num + 1)) {
                cur_num += 1;
                length += 1;
            }

            if length > result {
                result = length;
            }
        }
    });

    result
}
