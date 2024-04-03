/// 分割等和子集
/// 给你一个 只包含正整数 的 非空 数组 nums 。
/// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
/// 0-1 背包问题
pub fn can_partition(nums: Vec<i32>) -> bool {
    let weight: i32 = nums.iter().sum();

    // 奇数
    if weight & 1 == 1 {
        return false;
    }

    let weight = weight >> 1;

    let mut dp = vec![0; weight as usize];

    for i in 0..nums.len() {
        for j in (nums[i]..=weight).rev() {
            let j = j as usize;
            dp[j] = dp[j].max(dp[j - nums[i] as usize] + nums[i]);
            // 此处优化，提前返回
            if dp[j] == weight {
                return true;
            }
        }
    }

    dp[weight as usize] == weight
}
