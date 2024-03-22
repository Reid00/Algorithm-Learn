/// 只出现一次的数字  除了某个元素只出现一次以外，其余每个元素均出现两次
pub fn single_number(nums: Vec<i32>) -> i32 {
    // 0^任意值=任意值
    let mut res = 0;

    for v in nums {
        res ^= v;
    }

    res
}
