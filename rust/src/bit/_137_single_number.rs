// 除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
pub fn single_number(nums: Vec<i32>) -> i32 {
    // 如果某个元素出现了三次，那么将其二进制bit位，按位想加，则其和应该是3的倍数
    // 思路: 1. 按位想加
    // 2. 对3 取模
    // 3. 得到的元素就是出现一次的元素的二进制bnui
    let mut res = 0;

    // 32 位元素，循环32次，计算按位想加的和
    for i in 0..32 {
        let mut sum = 0;

        for v in nums.iter() {
            // v &1 ： 获取 v 的最低位（最右面bit 位的值）
            // 第 i 次计算，需要将 v >> i ， 右移i位，获取第i位的值
            sum += (v >> i) & 1
        }
        // 将sum 取模的值，左移i位，放到res i 位的位置上
        res |= (sum % 3) << i
    }

    res
}