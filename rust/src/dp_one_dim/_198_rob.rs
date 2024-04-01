/// 打家劫舍
pub fn rob(nums: Vec<i32>) -> i32 {
    let (mut prev, mut cur) = (0, 0);

    for v in nums {
        (prev, cur) = (cur, cur.max(v + prev));
    }

    cur
}
